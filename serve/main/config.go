package main

import "net/http"

var (
	fileSystem http.Handler
)

func init() {
	fileSystem = http.FileServer(http.Dir("/resources/"))
	//fileSystem = http.FileServer(http.Dir("D:\\Git\\intergrations\\alumni-association\\serve\\resources\\"))

}
