package memberCenter

import (
	"backend/public"
	"backend/service"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//信息展示 用于卡片和信息修改页面的信息展示
func Showmsg(w http.ResponseWriter, r *http.Request) {
	dbConn := service.GetDatabaseConnection()
	var err error
	//合并后要使用的代码 从前端的缓存中获取到id 通过id从数据库中拿到会员信息
	//错误信息返回
	reply := service.ReplyProto{
		Status: public.SUCCESS,
		Msg:    "success",
	}
	//getid, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reply.Status = -403 //r.Body 出错了
		reply.Msg = err.Error()
		return
	}
	//if getid == nil {
	//	reply.Status = -500 //传递的值为空
	//	reply.Msg = "invalid/nil request param"
	//	return
	//}
	//jsonMap := make(map[string]interface{})
	//err = json.Unmarshal(getid, &jsonMap)
	if err != nil {
		reply.Status = -400 //buf解码数据出现了错误
		reply.Msg = err.Error()
		return
	}

	//核心代码********************
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var id = session.Values["userID"].(int)
	//***********************

	//id := jsonMap["id"].(string)
	sqlStr := `select name,phone,email,status,sex,idcard,idcard_type,open_id from t_member where id = $1`
	//查询到的信息
	member_msg := service.Member_msg{
		Status: public.SUCCESS,
		Msg:    "success",
	}
	member_msg.Id = id
	err = dbConn.QueryRow(context.Background(), sqlStr, id).Scan(&member_msg.Name,
		&member_msg.Phone, &member_msg.Email,
		&member_msg.UserStatus, &member_msg.Sex,
		&member_msg.IdCard, &member_msg.IdCardType,
		&member_msg.Nickname)
	//fmt.Println(member_msg)
	if err != nil {
		reply.Status = -1003
		reply.Msg = err.Error()
		fmt.Println(err.Error())
		service.Resp(w, &reply)
		return
	}

	//resp(w, &msg)
	buf, err := json.Marshal(&member_msg)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code":-300,"msg":"%s"}`, err.Error())))
		fmt.Println(err.Error())
		return
	}
	w.Write(buf)

}
func GetOpinion(w http.ResponseWriter, r *http.Request) {
	dbConn := service.GetDatabaseConnection()
	var err error
	//合并后要使用的代码 从前端的缓存中获取到id 通过id从数据库中拿到会员信息
	//错误信息返回
	reply := service.ReplyProto{
		Status: public.SUCCESS,
		Msg:    "success",
	}
	//getid, err := ioutil.ReadAll(r.Body)
	//fmt.Println(getid)
	if err != nil {
		reply.Status = -403 //r.Body 出错了
		reply.Msg = err.Error()
		return
	}
	//if getid == nil {
	//	reply.Status = -500 //传递的值为空
	//	reply.Msg = "invalid/nil request param"
	//	return
	//}
	//jsonMap := make(map[string]interface{})
	//err = json.Unmarshal(getid, &jsonMap)
	if err != nil {
		reply.Status = -400 //buf解码数据出现了错误
		reply.Msg = err.Error()
		return
	}

	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var id = session.Values["userID"].(int)
	//id := jsonMap["id"].(string)
	//fmt.Println(id)
	sqlStr := `select status from t_member where id = $1`
	//var msg member_msg
	//查询到的信息
	member_msg := service.Member_msg{
		Status: public.SUCCESS,
		Msg:    "success",
	}
	member_msg.Id = id
	err = dbConn.QueryRow(context.Background(), sqlStr, id).Scan(&member_msg.UserStatus)
	if err != nil {
		reply.Status = -1003
		reply.Msg = err.Error()
		fmt.Println(err.Error())
		service.Resp(w, &reply)
		return
	}

	buf, err := json.Marshal(&member_msg)
	if err != nil {
		w.Write([]byte(fmt.Sprintf(`{"code":-300,"msg":"%s"}`, err.Error())))
		fmt.Println(err.Error())
		return
	}
	w.Write(buf)
}

//只改返回前端的
func Changemsg(w http.ResponseWriter, r *http.Request) {
	dbConn := service.GetDatabaseConnection()
	reply := service.ReplyProto{
		Status: public.SUCCESS,
		Msg:    "success",
	}

	newmsg, err := ioutil.ReadAll(r.Body)
	if err != nil {
		reply.Status = -400 //buf解码数据出现了错误
		reply.Msg = err.Error()
		service.Resp(w, &reply)
		panic(err.Error())
		return
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(newmsg, &jsonMap)
	//id := jsonMap["id"]
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var id = session.Values["userID"].(int)
	status_string := jsonMap["user_status"]
	phone := jsonMap["phone"]
	email := jsonMap["email"]
	idcard := jsonMap["idcard"]
	idcardtype := jsonMap["idcardtype"]
	name := jsonMap["name"]
	sex := jsonMap["sex"]
	sqlStr := ""
	status, _ := strconv.ParseInt(fmt.Sprint(status_string), 10, 64)
	if status == 1 {
		sqlStr = `update t_member set phone =$1,email = $2 , sex = $3 where id = $4`
		_, err = dbConn.Exec(context.Background(), sqlStr, phone, email, sex, id)
	} else {
		if status == 2 {
			status = 3
		}
		sqlStr = `update t_member set phone =$1,email = $2 ,idcard = $3 , idcard_type =$4, sex = $5,status = $6,name= $7 where id = $8`
		_, err = dbConn.Exec(context.Background(), sqlStr, phone, email, idcard, idcardtype, sex, status, name, id)
	}

	//合并后可以获取通过缓存获取id

	if err != nil {
		reply.Status = -300
		reply.Msg = "update information failed "
		service.Resp(w, &reply)
		panic(err.Error())
	}
	service.Resp(w, &reply)
	return
}
