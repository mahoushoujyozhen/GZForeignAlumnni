package activity

import (
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetQrcode(w http.ResponseWriter, r *http.Request) {
	var err error
	msg := Qrcode{
		Status: 0,
		Msg:    "success",
	}

	if strings.ToLower(r.Method) != "post" {
		msg.Status = -500
		msg.Msg = "invalid request,should be post"
		fmt.Println(msg.Status, msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -501
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	if buf == nil {
		msg.Status = -502
		msg.Msg = "invalid/nil request param"
		fmt.Println(msg.Status, msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -503
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	a := jsonMap["act_id"]
	if a == "" {
		msg.Status = -505
		msg.Msg = "invalid/act_id"
		fmt.Println(msg.Status, msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	actId, err := strconv.Atoi(a.(string))
	if err != nil {
		msg.Status = -506
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	url := fmt.Sprintf("http://mac.gzhu.edu.cn/alumni/#/login?act_id=%d", actId)
	// url := fmt.Sprintf("http://localhost:8080/#/userLogin?act_id=%d", actId)
	fmt.Println(url)
	var png []byte
	png, err = qrcode.Encode(url, qrcode.Medium, 256)
	if err != nil {
		msg.Status = -408
		msg.Msg = err.Error()
		fmt.Println(msg.Msg)
		RespQrcode(w, &msg)
		return
	}

	msg.Qrcode = png
	RespQrcode(w, &msg)
}

func RespQrcode(w http.ResponseWriter, msg *Qrcode) {
	if w == nil || msg == nil {
		w.Write([]byte(fmt.Sprintf(`{"status":-447,"msg":"%s"}`,
			"call respErr with invalid w/msg")))
		fmt.Println("call respErr with invalid w/msg")
		return
	}

	buf, err := json.Marshal(&msg)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"status":-448,"msg":"%s"}`, err.Error())))
		fmt.Println(err.Error())
		return
	}

	_, err = w.Write(buf)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"status":-449,"msg":"%s"}`, err.Error())))
		fmt.Println(err.Error())
	}
}
