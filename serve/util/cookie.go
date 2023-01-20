package util

import "github.com/gorilla/sessions"

var cookie *sessions.CookieStore

func init() {
	cookie = sessions.NewCookieStore([]byte("sessionID"))
}

func GetStore() *sessions.CookieStore {
	return cookie
}
