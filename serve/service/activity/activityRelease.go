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

func Act_release(w http.ResponseWriter, r *http.Request) {

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

	//err=r.ParseForm()
	//if err!=nil{
	//	println(1)
	//	msg.Status = -1030
	//	msg.Msg = err.Error()
	//	fmt.Println(msg.Msg)
	//	service.Resp(w, &msg)
	//	return
	//}
	//name :=  r.FormValue("name")
	//address := r.FormValue("address")
	//start_Time :=  r.FormValue("start_Time")
	//end_Time :=  r.FormValue("end_Time")
	//profile :=  r.FormValue("profile")
	//img_url :=  r.FormValue("img_url")
	//fmt.Println(name)

	fmt.Printf("活动发布：name:%s,address:%s,start_Time:%s,end_Time:%s,profile:%s,img_url:%s",
		name, address, start_Time, end_Time, profile, img_url)

	if name == "" || address == "" || start_Time == "" || end_Time == "" || profile == "" || img_url == "" {
		msg.Status = -500
		msg.Msg = "invalid/empty input"
		fmt.Println(msg.Msg)
		return
	}

	s := `insert into t_activity_module (act_name,act_address,start_time,end_time,act_profile,img_url,apply_num) 
			values($1,$2,$3,$4,$5,$6,0)`
	_, err = dbConn.Exec(context.Background(), s, name, address, start_Time, end_Time, profile, img_url)
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
