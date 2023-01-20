package activity

import (
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func SigninList(w http.ResponseWriter, r *http.Request) {
	//接收结构
	type MsgRequest struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

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

	if buf == nil {
		msg.Status = -500
		msg.Msg = "invalid/nil request param"
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -400
		msg.Msg = err.Error()
		return
	}
	idTemp := jsonMap["act_id"].(string)
	act_id, err := strconv.Atoi(idTemp)
	if err != nil {
		msg.Status = -402
		msg.Msg = err.Error()
		return
	}

	s := fmt.Sprintf(`select signin_info from t_activity_module where act_id=$1`)
	rs, err := dbConn.Query(context.Background(), s, act_id)
	if err != nil {
		fmt.Println(err)
		return
	}
	var name []MsgRequest
	for rs.Next() {
		err = rs.Scan(&name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	var users []string
	for i := 0; i < len(name); i++ {
		users = append(users, fmt.Sprintf(`{"name":"%s","phone":"%s"}`,
			name[i].Name, name[i].Phone))
	}
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(users, ",")))
	service.Resp(w, &msg)
}
