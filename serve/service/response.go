package service

import (
	"backend/logging"
	"encoding/json"
	"fmt"
	"net/http"
)

func Resp(w http.ResponseWriter, msg *ReplyProto) {

	if w == nil || msg == nil {
		fmt.Println("call respErr with invalid w/msg")
		return
	}

	//将msg结构体转化为json字符串
	buf, err := json.Marshal(&msg)

	fmt.Println(msg)

	fmt.Println("buf---------",buf)
	//生成json字符串失败时
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code":-300,"msg":"%s"}`, err.Error())))
		fmt.Println(err.Error())
		return
	}

	w.Write(buf)
	//fmt.Println(string(buf))
}

func ResponseService(responseWriter http.ResponseWriter, responseBean *ResponseBean) {
	var errors error
	// 设置响应头
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	responseWriter.Header().Set("Content-Type", "application/json;charset=utf-8")
	// 设置响应体
	var response []byte
	response, errors = json.Marshal(&responseBean)
	if errors != nil {
		logging.Error(errors.Error())
	}
	responseWriter.WriteHeader(http.StatusOK)
	// 返回客户端
	_, errors = responseWriter.Write(response)
	if errors != nil {
		logging.Error(errors.Error())
	}
}
