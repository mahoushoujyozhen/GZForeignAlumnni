package activity

import (
	"backend/service"
	"backend/util"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//取消活动
func CancelActivity(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	if strings.ToLower(r.Method) != "delete" {
		msg.Status = -400
		msg.Msg = "invalid request,should be delete"
		service.Resp(w, &msg)
		return
	}
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	memberid := session.Values["userID"].(int)

	//memberid := r.URL.Query().Get("user_id")
	actid := r.URL.Query().Get("act_id")

	if actid == "" {
		msg.Status = -1010
		msg.Msg = "missing user_id or act_id"
		service.Resp(w, &msg)
		return
	}
	//_, err := strconv.ParseInt(memberid, 10, 64)
	_, err := strconv.ParseInt(actid, 10, 64)
	if err != nil {
		msg.Msg = "member_id or act_id should be number"
		msg.Status = -1020
		service.Resp(w, &msg)
		return
	}

	s := fmt.Sprintf(`delete from t_join where member_id=$1 and act_id=$2`)
	_, err = dbConn.Exec(context.Background(), s, memberid, actid)
	if err != nil {
		fmt.Println("33333")
		msg.Status = -1000
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		return

	}

	fmt.Println("activity had canceled")

	ss := `select apply_num from t_activity_module where act_id=$1`
	rs := dbConn.QueryRow(context.Background(), ss, actid) //读取每一个活动id
	//if err != nil {
	//	fmt.Println("12341345")
	//	msg.Status = -1000
	//	msg.Msg = err.Error()
	//	service.Resp(w, &msg)
	//	return
	//}
	var applynum int64
	err = rs.Scan(&applynum)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	applynum -= 1
	s = `update t_activity_module set apply_num=$1 where act_id=$2 `
	_, err = dbConn.Exec(context.Background(), s, applynum, actid)
	if err != nil {
		fmt.Println("Exec err=", err)
		return
	}
	fmt.Println("data updated successfully")
	msg.Data = []byte(fmt.Sprintf(`{"act_id":%s,"apply_num":%d}`, actid, applynum))
	service.Resp(w, &msg)
}

//活动报名或取消报名
func ApplyActivity(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}

	if strings.ToLower(r.Method) != "put" {
		msg.Status = -400
		msg.Msg = "invalid request,should be put"
		service.Resp(w, &msg)
		return

	}

	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var memberid = session.Values["userID"].(int)
	//memberid := r.URL.Query().Get("user_id")
	actid := r.URL.Query().Get("act_id")

	if actid == "" {
		msg.Status = -1010
		msg.Msg = "missing user_id or act_id"
		service.Resp(w, &msg)
		return

	}

	//memberId, err := strconv.ParseInt(memberid, 10, 64)
	actId, err := strconv.ParseInt(actid, 10, 64)
	if err != nil {
		msg.Msg = "member_id or act_id should be number"
		msg.Status = -1020
		service.Resp(w, &msg)
		return

	}

	s := `insert into t_join (member_id, act_id) values ($1, $2)`
	_, err = dbConn.Exec(context.Background(), s, memberid, actId)
	if err != nil {
		msg.Status = -1000
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		return

	}

	fmt.Println("activity had applied")

	s = `select apply_num from t_activity_module where act_id=$1`
	rs := dbConn.QueryRow(context.Background(), s, actid) //读取每一个活动id
	//if err != nil {
	//	fmt.Println("12341345")
	//	msg.Status = -1000
	//	msg.Msg = err.Error()
	//	service.Resp(w, &msg)
	//	return
	//}
	var applynum int64
	err = rs.Scan(&applynum)
	if err != nil {
		fmt.Println(err.Error())
		return

	}
	applynum += 1
	s = `update t_activity_module set apply_num=$1 where act_id=$2 `
	_, err = dbConn.Exec(context.Background(), s, applynum, actid)
	if err != nil {
		fmt.Println("Exec err=", err)
		return

	}

	fmt.Println("data updated successfully")
	msg.Data = []byte(fmt.Sprintf(`{"act_id":%s,"apply_num":%d}`, actid, applynum))
	service.Resp(w, &msg)
}
