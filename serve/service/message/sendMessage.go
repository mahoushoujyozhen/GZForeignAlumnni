package message

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

//切换页面时列出用户
func ListUser(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	if strings.ToLower(r.Method) != "post" {
		msg.Status = -5000
		msg.Msg = fmt.Sprintf("invalid request，should be post")
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -5001
		msg.Msg = fmt.Sprintf(err.Error() + "Error ReadAll from r.body")
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -5002
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	currentPage := jsonMap["userCurrentPage"]
	equipment := jsonMap["equipment"]
	var num int64
	var s string
	if equipment == "PC" {
		num = 11
		s = `select id , name ,count(*) over() as total from t_member limit 11 offset ($1 - 1)*11;`
	} else {
		num = 9
		s = `select id , name ,count(*) over() as total from t_member limit 9 offset ($1 - 1)*9;`
	}
	rows, err := dbConn.Query(context.Background(), s, currentPage)
	defer rows.Close()
	if err != nil {
		//连接查询出错
		msg.Status = -5003
		msg.Msg = err.Error()
		logging.Error(fmt.Sprintf(msg.Msg + "Error:Query in listuser"))
		service.Resp(w, &msg)
		return
	}

	var user []string
	var name string
	var id int64
	var count int64
	for rows.Next() {
		err = rows.Scan(&id, &name, &count)
		if err != nil {
			msg.Status = -5003
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
			return
		}

		if count%num == 0 {
			count = count / num
		} else {
			count = count/num + 1
		}
		user = append(user, fmt.Sprintf(`{"id":%d,"name":"%s"}`, id, name))
	}
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(user, ",")))
	msg.PageCount = count
	w.Header().Set("Content_Type", "application/json;charset=utf-8")
	service.Resp(w, &msg)
}

func SearchUser(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	if strings.ToLower(r.Method) != "post" {
		msg.Status = -5000
		msg.Msg = fmt.Sprintf("invalid request, should be post")
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -5001
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -5002
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	searchContent := jsonMap["searchContent"]
	currentPage := jsonMap["userCurrentPage"]
	equipment := jsonMap["equipment"]
	searchUser := fmt.Sprint(searchContent)
	searchUser = "%" + searchUser + "%"
	var name string
	var id int64
	//allUsers := make([]Receiver, 0)
	var allUser []string
	s := `select id,name from t_member where name LIKE $1`
	rc, err := dbConn.Query(context.Background(), s, searchUser)
	defer rc.Close()
	for rc.Next() {
		err = rc.Scan(&id, &name)
		if err != nil {
			msg.Status = -5003
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
			return
		}

		allUser = append(allUser, fmt.Sprintf(`{"id":%d,"name":"%s"}`, id, name))
	}
	var num int64
	if equipment == "PC" {
		num = 11
		s = `select id , name ,count(*) over() as total from t_member where name LIKE $1 limit 11 offset ($2 -1) * 11`
	} else {
		num = 9
		s = `select id , name ,count(*) over() as total from t_member where name LIKE $1 limit 9 offset ($2 -1) * 9`
	}
	//s = `select id , name ,count(*) over() as total from t_member where name LIKE $1 limit 9 offset ($2 -1) * 9`
	rows, err := dbConn.Query(context.Background(), s, searchUser, currentPage)
	defer rows.Close()
	if err != nil {
		msg.Status = -5004
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	var user []string
	var Count int64
	for rows.Next() {
		err = rows.Scan(&id, &name, &Count)
		if err != nil {
			msg.Status = -5005
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
			return
		}
		
		if Count%num == 0 {
			Count = Count / num
		} else {
			Count = Count/num + 1
		}
		user = append(user, fmt.Sprintf(`{"id":%d,"name":"%s"}`, id, name))
	}
	//buf, _ = json.Marshal(user)
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(user, ",")))
	msg.AllUser = []byte(fmt.Sprintf(`[%s]`, strings.Join(allUser, ",")))
	msg.PageCount = Count
	w.Header().Set("Content_Type", "application/json;charset=utf-8")
	service.Resp(w, &msg)
}

//将选择的用户名加载到输入框中
func ListReceiver(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	if strings.ToLower(r.Method) != "post" {
		msg.Status = -5000
		msg.Msg = fmt.Sprintf("invalid request,should be post")
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return

	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -5002
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return

	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -5003
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return

	}
	id := jsonMap["searchId"]
	Uid := fmt.Sprint(id)
	sep := ";"
	userId := strings.Split(Uid, sep)
	var receiver []string
	//var  []string
	var name string
	s := `select name from t_member where id=$1`
	for i := range userId {
		ID, _ := strconv.ParseInt(userId[i], 10, 64)
		//fmt.Println(id1)
		rc := dbConn.QueryRow(context.Background(), s, ID)
		err = rc.Scan(&name)
		if err != nil {
			msg.Status = -5004
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
		}
		receiver = append(receiver, fmt.Sprintf(`{"id":%d,"name":"%s"}`, ID, name))
	}
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(receiver, ",")))
	service.Resp(w, &msg)
}

//消息发送界面打开后，调用函数加载全部用户
func ListAllUser(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	if strings.ToLower(r.Method) != "get" {
		msg.Status = -5000
		msg.Msg = fmt.Sprintf("invalid request,should be get")
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	s := `select id ,name from t_member;`
	rows, err := dbConn.Query(context.Background(), s)
	defer rows.Close()
	if err != nil {
		msg.Status = -5001
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	var allUser []string
	var name string
	var id int64
	for rows.Next() {
		err = rows.Scan(&id, &name)
		if err != nil {
			msg.Status = -5002
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
			return
		}

		allUser = append(allUser, fmt.Sprintf(`{"id":%d,"name":"%s"}`, id, name))
	}
	msg.AllUser = []byte(fmt.Sprintf(`[%s]`, strings.Join(allUser, ",")))
	w.Header().Set("Content_Type", "application/json;charset=utf-8")
	service.Resp(w, &msg)
}
