package main

import (
	"backend/service"
	"backend/util"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
)

func BaseServer(responseWriter http.ResponseWriter, request *http.Request) {
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}
	store := util.GetStore()
	session, _ := store.Get(request, "sessionID")
	authenticated := session.Values["authenticated"]
	fmt.Println("authenticated:", authenticated)
	if authenticated == nil || authenticated == false {
		msg.Status = -999
		msg.Msg = "登录认证失败！"
		fmt.Println("诺曼底登陆！！！")
		service.Resp(responseWriter, &msg)
		return
	}

	responseWriter.Header().Set("Location", "http://172.22.72.229:9090/")

	//解决跨域问题
	responseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	// 获取请求URI
	uri := request.URL.RequestURI()
	// 编译正则表达式
	cRegx := regexp.MustCompile(routeRegx)
	// 分解请求URI
	groups := cRegx.FindStringSubmatch(uri)
	// 检查请求格式和方法
	requestType := groups[1]
	requestTarget := groups[2]
	requestAction := groups[3]
	handleRoute := fmt.Sprintf("/%v/%v", requestTarget, requestAction)
	if len(requestType) == 0 || len(requestTarget) == 0 || len(requestAction) == 0 || routeMap[handleRoute] == nil {
		// 不合法请求
		msg.Status = -499
		msg.Msg = "请求参数不合法"
		service.Resp(responseWriter, &msg)
		return
	}
	isAdmin := session.Values["isAdmin"]
	fmt.Println("isAdmin:", isAdmin)
	if requestTarget == "admin" && (isAdmin == nil || isAdmin == false) {
		msg.Status = -1999
		msg.Msg = "管理员账号未登录！"
		service.Resp(responseWriter, &msg)
		return
	}

	// 根据请求路径分发
	handler := reflect.ValueOf(routeMap[handleRoute])
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(responseWriter)
	inputs[1] = reflect.ValueOf(request)
	handler.Call(inputs)

}
