package main

import (
	"fmt"
	"net/http"
	"os"
	"path"
	"strings"
)

var docRoot = "/var/deploy/alumni/f"

func staticFileServe() (fs http.Handler) {
	docRoot = path.Clean(docRoot)
	_, err := os.Stat(docRoot)
	if err != nil {
		fmt.Println(fmt.Sprintf("配置项docRoot:%s不存在或无权限，请修正", docRoot))
		os.Exit(-1)
		return
	}

	idxFile := docRoot + "/index.html"
	_, err = os.Stat(idxFile)
	if os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf("配置项docRoot:%s不存在，请修正", idxFile))
		os.Exit(-1)
		return
	}

	rootFS := http.FileServer(http.Dir(docRoot))
	fs = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if w == nil || r == nil {
			fmt.Println("w or r is nil", idxFile)
			os.Exit(-1)
			return
		}
		rPath := r.URL.Path
		if !strings.HasPrefix(rPath, "/") {
			rPath = "/" + rPath
			r.URL.Path = rPath
		}

		dstFileName := path.Clean(docRoot + rPath)
		_, err = os.Stat(dstFileName)
		if os.IsNotExist(err) {
			r.URL.Path = "/index.html"
		}

		rootFS.ServeHTTP(w, r)
	})
	return fs
}
