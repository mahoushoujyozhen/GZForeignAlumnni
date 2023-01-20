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

//发布活动

func Act_change(w http.ResponseWriter, r *http.Request) {

	fmt.Println("修改活动")
	var err error
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}

	if strings.ToLower(r.Method) != "post" {
		msg.Status = -400
		msg.Msg = "invalid request,should be post"
		fmt.Println(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -403
		msg.Msg = err.Error()
		fmt.Println(msg.Msg)
		return
	}

	if buf == nil {
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
	name := jsonMap["name"]
	address := jsonMap["address"]
	start_Time := jsonMap["start_Time"]
	end_Time := jsonMap["end_Time"]
	profile := jsonMap["profile"]
	img_url := jsonMap["img_url"]
	cur_actId := jsonMap["cur_actId"]
	fmt.Printf("活动发布：name:%s,address:%s,start_Time:%s,end_Time:%s,profile:%s,img_url:%s",
		name, address, start_Time, end_Time, profile, img_url)

	if name == "" || address == "" || start_Time == "" || end_Time == "" || profile == "" || img_url == "" {
		msg.Status = -500
		msg.Msg = "invalid/empty input"
		fmt.Println(msg.Msg)
		return
	}

	s := `update t_activity_module set act_name=$1,act_address=$2,start_time=$3,end_time=$4,act_profile=$5,img_url=$6 where act_id=$7`
	_, err = dbConn.Exec(context.Background(), s, name, address, start_Time, end_Time, profile, img_url, cur_actId)
	if err != nil {
		msg.Status = -1030
		msg.Msg = err.Error()
		fmt.Println(msg.Msg)
		service.Resp(w, &msg)
		return
	}
	msg.Data = []byte(fmt.Sprintf(`{"name":"%s","address":"%s","start_Time":"%s","end_Time":"%s","profile":"%s","img_url":"%s"}`, name, address, start_Time, end_Time, profile, img_url))
	service.Resp(w, &msg)

	return

}
