package activity

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Act_delete(w http.ResponseWriter, r *http.Request) {
	fmt.Println("删除活动")

	var err error
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}

	//strings.ToLower()函数接收一个字符串，并将字符串中的每个ASCII字符转换成大写，最后将新的字符串返回。
	if strings.ToLower(r.Method) != "post" {
		fmt.Println("2")
		msg.Status = -400
		msg.Msg = "invalid request,should be post"
		fmt.Println(msg.Msg)
		service.Resp(w, &msg) //将返回信息msg传到前端
		return
	}

	//从形参中读取数据直到发生错误或遇到EOF，返回读到的数据。
	buf, err := ioutil.ReadAll(r.Body) //把前端发送的body(请求体) 内容读入字符串buf
	if err != nil {
		fmt.Println("3")
		msg.Status = -403
		msg.Msg = err.Error() //err.Error()作用是返回错误信息
		fmt.Println(msg.Msg)
		return
	}

	if buf == nil {
		fmt.Println("4")
		msg.Status = -500
		msg.Msg = "invalid/nil request param"
		fmt.Println(msg.Msg)
		return
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -400
		msg.Msg = err.Error()
		fmt.Println(msg.Msg)
		return
	}
	id := jsonMap["act_id"] //前端传过来的id
	fmt.Println("id:", id)

	s1 := `delete from t_join where act_id = $1`
	s2 := `delete from t_activity_module where act_id = $1`

	_, err = dbConn.Exec(context.Background(), s1, id)
	if err != nil {
		msg.Status = -1030
		msg.Msg = err.Error()
		fmt.Println(msg.Msg)
		service.Resp(w, &msg)
		return
	}
	_, err = dbConn.Exec(context.Background(), s2, id)
	if err != nil {
		msg.Status = -1030
		msg.Msg = err.Error()
		fmt.Println(msg.Msg)
		service.Resp(w, &msg)
		return
	}
}
