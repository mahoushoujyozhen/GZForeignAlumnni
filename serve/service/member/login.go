package member

import (
	"backend/logging"
	"backend/service"
	"backend/util"
	"encoding/json"
	"fmt"
	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v4"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	//解决跨域问题
	w.Header().Set("Access-Control-Allow-Origin", "*")
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}

	buf, err := ioutil.ReadAll(r.Body)
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		logging.Error(err.Error())
		msg.Status = -501
		msg.Msg = "json unmarshal error"
		service.Resp(w,&msg)
	}

	username := jsonMap["username"]
	password := jsonMap["password"]

	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")

	if username == "" || password == "" {
		msg.Status = -500
		msg.Msg = "invalid/empty username/password"
		service.Resp(w, &msg)
		return
	}

	s := `select id,open_id,status from t_member where open_id = $1 and password=crypt($2,password) `

	var userID , status int
	result := dbConn.QueryRow(context.Background(), s, username, password)
	err = result.Scan(&userID, &username,&status)
	nonexistent := err == pgx.ErrNoRows

	if nonexistent {
		msg.Status = -20004
		msg.Msg = err.Error()
		logging.Error(err)
		service.Resp(w, &msg)
		return
	}


	//保存session内容
	session.Values["authenticated"] = true
	session.Values["userID"] = userID
	session.Values["username"] = username

	//status为4，说明其为管理员，管理员认证成功
	if status == 4 {
		session.Values["isAdmin"] = true
	}

	//登录成功保存会话
	err = sessions.Save(r, w)

	if err != nil {
		msg.Status = -502
		msg.Msg = err.Error()
		logging.Error(err)
		return
	}

	msg.Data = []byte(fmt.Sprintf(`
   {"userID":%d,"username":"%s"}`, userID, username))
	service.Resp(w, &msg)
	return
}

func Logout(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	session.Values["authenticated"] = false
	session.Values["isAdmin"] = false
	//删除cookie
	session.Options.MaxAge = -1
	err := session.Save(r, w)
	if err != nil {
		msg.Status = -502
		msg.Msg = err.Error()
		logging.Error(err)
		return
	}
	fmt.Println(w, "Successfully Logged Out")
	service.Resp(w,&msg)
	return
}

func IsLogin(w http.ResponseWriter, r *http.Request){
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")

	isAdmin := session.Values["isAdmin"]
	if isAdmin == true{
		msg.Status = 201
		msg.Msg = "管理员账号已登录"
		service.Resp(w,&msg)
		return
	}

	isUser := session.Values["authenticated"]
	if isUser == true {
		msg.Status = 202
		msg.Msg = "用户账号已登录"
		service.Resp(w,&msg)
		return
	}

	msg.Msg = "账号未登陆"
	service.Resp(w,&msg)
	return
}
