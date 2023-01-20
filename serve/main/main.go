package main

import (
	"backend/logging"
	"backend/public"
	"backend/service/file"
	"backend/service/member"
	"fmt"
	"net/http"
)

func main() {
	var errors error

	// HTTP请求
	http.HandleFunc("/api/file/generateId", file.GenerateIdService)
	http.HandleFunc("/api/file/upload", file.UploadService)
	http.HandleFunc("/api/file/delete", file.DeleteService)
	http.HandleFunc("/api/", BaseServer)
	http.HandleFunc("/api/user/is_login", member.IsLogin)
	http.HandleFunc("/api/user/login", member.Login)
	http.HandleFunc("/api/user/submit", member.Submit)
	// 静态资源请求

	http.Handle("/resources/", http.StripPrefix("/resources/", fileSystem))
	http.Handle("/", staticFileServe())

	errors = http.ListenAndServe(":9090", nil)
	if errors != nil {
		logging.Error(public.ERR_SERVER_START)
	} else {
		fmt.Println(public.MSG_SERVER_START)
		logging.Error(public.MSG_SERVER_START)
	}

}
