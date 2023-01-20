package message

import (
	"backend/logging"
	e2 "backend/public"
	"backend/service"
	"backend/service/file"
	"backend/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

func MessageService(responseWriter http.ResponseWriter, request *http.Request) {
	// 获取请求参数
	store := util.GetStore()
	session, _ := store.Get(request, "sessionID")
	userID := session.Values["userID"].(int)
	id := fmt.Sprintf("%d", userID)
	if len(id) == 0 {
		service.ResponseService(responseWriter, service.ResponseBeanFactory(e2.ERROR, e2.GetMsg(e2.INVALID_PARAMS), nil))
	} else {
		SocketService(id, hub, responseWriter, request)
	}
}

// SocketService 真正的服务函数
func SocketService(id string, hub *Hub, responseWriter http.ResponseWriter, request *http.Request) {
	var errors error
	// 将普通的HTTP请求升级到长连接WebSocket
	var connection *websocket.Conn
	connection, errors = upgrader.Upgrade(responseWriter, request, responseWriter.Header())
	if errors != nil {
		return
	}

	// 注册客户端
	client := &Client{id: id, hub: hub, connection: connection, send: make(chan []byte, 512)}
	client.hub.register <- client

	// 开启双向读写
	go client.ReadPump()
	go client.WritePump()

	//如果非管理员连接
	if id != "-1" {
		// 根据id查询数据库是否存在离线时的消息
		var notifications []service.NotificationBean
		notifications, errors = file.ReadMessageByUserId(id)
		if errors != nil {
			logging.Error(errors.Error())
			return
		}
		client.hub.send <- Message{ClientId: id, Details: notifications}

		// 删除离线消息(如果有错误自动回滚)
		errors = file.DeleteMessageByUserId(id)
		if errors != nil {
			logging.Error(errors.Error())
		}
	}
}

func HubFactory() *Hub {
	return &Hub{clients: make(map[string]*Client), send: make(chan Message), register: make(chan *Client), unregister: make(chan *Client)}
}

func (hub *Hub) Run() {
	for {
		select {
		case client := <-hub.register: // 用户上线,连接客户端
			hub.clients[client.id] = client
		case client := <-hub.unregister: // 用户下线,移除客户端
			if _, ok := hub.clients[client.id]; ok {
				delete(hub.clients, client.id)
				close(client.send)
			}
		case message := <-hub.send:
			if client, ok := hub.clients[message.ClientId]; ok {
				// 向指定客户端传输消息(懒得处理错误)
				dataString, _ := json.Marshal(&message)
				select {
				case client.send <- dataString:
				default:
					delete(hub.clients, client.id)
					close(client.send)
				}
			}
		}
	}
}

// ReadPump 读客户端,也是是否下线的依据
func (client *Client) ReadPump() {
	var errors error
	defer func() {
		logging.Error(errors)
		client.hub.unregister <- client
		client.connection.Close()
	}()

	client.connection.SetReadLimit(MAX_READ_SIZE)
	client.connection.SetReadDeadline(time.Now().Add(PONG_WAIT))
	client.connection.SetPongHandler(func(string) error {
		client.connection.SetReadDeadline(time.Now().Add(PONG_WAIT))
		return nil
	})
	for {
		// 第二个参数是客户端发来的信息,在本次项目中只有管理员的信息可以接收
		var message []byte
		_, message, errors = client.connection.ReadMessage()
		if errors != nil {
			if websocket.IsUnexpectedCloseError(errors, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Println(errors)
				return
			}
			break
		}
		ids, title, sendTime, content, isValid := GetMessageParams(message)
		//fmt.Println(ids, title, sendTime)
		//fmt.Println(title)
		if !isValid {
			logging.Error(errors.Error())
		} else {
			DistributeMessage(ids, title, sendTime, content)
		}
	}
}

// WritePump 写客户端
func (client *Client) WritePump() {
	// 设置写延迟
	ticker := time.NewTimer(PING_WAIT)
	defer func() {
		ticker.Stop()
		client.connection.Close()
	}()
	for {
		select {
		case message, ok := <-client.send: // 发送消息
			client.connection.SetWriteDeadline(time.Now().Add(WRITE_DEADLINE))
			if !ok {
				client.connection.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			writer, err := client.connection.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			writer.Write(message)

			length := len(client.send)
			for i := 0; i < length; i++ {
				writer.Write(NEW_LINE)
				writer.Write(<-client.send)
			}
			if err := writer.Close(); err != nil {
				return
			}
		case <-ticker.C:
			client.connection.SetWriteDeadline(time.Now().Add(WRITE_DEADLINE))
			if err := client.connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// GetMessageParams 解析获取消息参数
func GetMessageParams(message []byte) ([]string, string, int64, string, bool) {
	var errors error
	// 解析JSON
	params := make(map[string]interface{})
	errors = json.Unmarshal(message, &params)
	if errors != nil {
		return nil, "", 0, "", false
	}
	// 校验参数
	// "Unmarshal"对数值类型的都会转化为"float64"
	//fmt.Println("params:", params)
	ids := params["ids"].([]interface{})
	title := params["title"].(string)
	sendTime := util.Float64ToInt64(params["time"].(float64))
	content := params["content"].(string)
	if len(ids) == 0 || len(title) == 0 || len(content) == 0 {
		return nil, "", 0, "", false
	}
	// 转化"string"数组
	var transIds []string
	for i := 0; i < len(ids); i++ {
		transIds = append(transIds, ids[i].(string))
	}
	return transIds, title, sendTime, content, true
}

// DistributeMessage 根据ids判断用户是否在线,如果在线,传输数据,否则写入数据库
func DistributeMessage(ids []string, title string, sendTime int64, content string) {
	var errors error
	// 遍历ids
	for i := 0; i < len(ids); i++ {
		client, isExist := hub.clients[ids[i]]
		if isExist {
			// 生成唯一Id
			var mid string
			mid, errors = util.GenerateUniqueId(snowflake)
			if errors != nil {
				logging.Error(errors.Error())
			} else {
				// 传输数据
				notifications := []service.NotificationBean{{Id: mid, Title: title, Time: sendTime, Content: content}}
				client.hub.send <- Message{ClientId: ids[i], Details: notifications}
			}
		} else {
			errors = file.InsertNewMessage(ids[i], title, sendTime, content)
			if errors != nil {
				logging.Error(errors.Error())
			}
		}
	}
}
