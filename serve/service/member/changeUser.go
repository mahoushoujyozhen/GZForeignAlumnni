package member

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ChangeUsers(w http.ResponseWriter, r *http.Request) {
	var err error

	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -403
		msg.Msg = err.Error()
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -400
		msg.Msg = err.Error()
		return
	}
	id := jsonMap["id"]
	sign := jsonMap["sign"]
	opinion := jsonMap["book"]
	fmt.Println(id)
	fmt.Println(sign)
	if sign == "1" {
		fmt.Println("数字1")
		s := fmt.Sprintf(`update t_member set status=$2,opinion=$3 where id = $1 `)
		fmt.Println(s)
		_, err = dbConn.Exec(context.Background(), s, id, 2, opinion)
		if err != nil {
			msg.Status = -1030
			msg.Msg = err.Error()
			logging.Error(msg.Msg) //错误时输出日志
			service.Resp(w, &msg)
			return
		}
		msg.Data = []byte(fmt.Sprintf(`{"id":%s}`, id))
		service.Resp(w, &msg)
	} else if sign == "2" {
		fmt.Println("数字2")
		s := fmt.Sprintf(`update t_member set status=$2,opinion=$3 where id = $1 `)
		fmt.Println(s)
		_, err = dbConn.Exec(context.Background(), s, id, 1, opinion)
		if err != nil {
			msg.Status = -1030
			msg.Msg = err.Error()
			logging.Error(msg.Msg) //错误时输出日志
			service.Resp(w, &msg)
			return
		}
		msg.Data = []byte(fmt.Sprintf(`{"id":%s}`, id))
		service.Resp(w, &msg)

	}

}
