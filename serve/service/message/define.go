package message

import (
	"backend/service"
	utils "backend/util"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	"net/http"
	_ "net/http"
	"time"
	_ "time"
)

var dbConn = service.GetDatabaseConnection()

type Receiver struct {
	//用户的id
	Id int64 `json:"id"`
	//用户姓名
	Name    string     `json:"name"`
	Count   int64      `json:"count"`
	Alluser []Receiver `json:"alluser"`
}

type MessageContent struct {
	//消息标题
	Title string `json:"title"`
	//消息正文
	Content string `json:"content"`
}

type Message1 struct {
	//用户ID
	ID int64 `json:"id"`
	//用户姓名
	Name string `json:"name"`
	//消息内容
	Data MessageContent `json:"data"`
}

type Client struct {
	id         string
	hub        *Hub
	connection *websocket.Conn
	send       chan []byte
}

type Hub struct {
	clients    map[string]*Client //客户端连接"池"
	send       chan Message
	register   chan *Client // 连接客户端
	unregister chan *Client // 断开客户端
}

type Message struct {
	ClientId string                     `json:"clientId,omitempty"`
	Details  []service.NotificationBean `json:"details,omitempty"`
}

const (
	WRITE_DEADLINE = time.Second * 10
	PONG_WAIT      = time.Second * 60 * 10 // 这个好像是设置自动断开事件的
	PING_WAIT      = (PONG_WAIT * 9) / 10
	MAX_READ_SIZE  = 4096
)

var (
	NEW_LINE  = []byte{'\n'}
	snowflake *utils.Snowflake
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  10240,
	WriteBufferSize: 10240,
	CheckOrigin: func(request *http.Request) bool {
		return true
	}, // 拦截跨域请求
}

var hub *Hub

func init() {
	hub = HubFactory()
	// 开启WebSocket服务端
	go hub.Run()
	snowflake, _ = utils.GetSnowflake(time.Now().UnixMilli(), 1, 1, 1)
}
