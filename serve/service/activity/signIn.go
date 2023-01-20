package activity

import (
	"backend/service"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"gopkg.in/guregu/null.v4"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	var err error
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}

	if strings.ToLower(r.Method) != "post" {
		msg.Status = -400
		msg.Msg = "invalid request,should be post"
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -401
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	if buf == nil {
		msg.Status = -402
		msg.Msg = "invalid/nil request param"
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -403
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	//u := jsonMap["user_id"]
	a := jsonMap["act_id"]
	if a == "" {
		msg.Status = -405
		msg.Msg = "invalid/empty name/phone/act_id"
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	actId, err := strconv.Atoi(a.(string))
	if err != nil {
		msg.Status = -407
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var userId = session.Values["userID"].(int)

	//查看用户是否报名
	var uID, aID int
	isApply := false
	s := `select member_id, act_id from t_join where member_id=$1 and act_id=$2`
	rc := dbConn.QueryRow(context.Background(), s, userId, actId)
	err = rc.Scan(&uID, &aID)
	if err != nil || (uID == 0 && aID == 0) {
		fmt.Println("用户未报名活动")
	} else {
		isApply = true
		fmt.Println("用户已报名活动")
	}

	if !isApply {
		msg.Data = []byte(fmt.Sprintf(`{"isApplied":%t}`, isApply))
		service.Resp(w, &msg)
		return
	}

	//查看活动是否已经结束
	var endTime null.String
	isOver := false
	s = `select end_time from t_activity_module where act_id=$1`
	rc = dbConn.QueryRow(context.Background(), s, actId)
	err = rc.Scan(&endTime)
	if err != nil {
		msg.Status = -408
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	//将活动结束时间与现在进行比较，返回bool，为true即活动结束；false表示活动未结束
	now := time.Now()
	local, _ := time.LoadLocation("Local")
	time, err := time.ParseInLocation("2006-01-02 15:04", endTime.String, local)
	if err != nil {
		msg.Status = -412
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	isOver = now.After(time)
	if isOver {
		fmt.Println("活动已结束")
		msg.Data = []byte(fmt.Sprintf(`{"isOver":%t}`, isOver))
		service.Resp(w, &msg)
		return
	}

	//如果已经报名，则进入签到信息录入
	s = fmt.Sprintf(`select name, phone from t_member where id = $1`)
	rc = dbConn.QueryRow(context.Background(), s, userId)
	var name null.String
	var phone null.String
	err = rc.Scan(&name, &phone)
	if err != nil {
		msg.Status = -409
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	s = fmt.Sprintf(`select signin_info from t_activity_module where act_id = $1`)
	rc = dbConn.QueryRow(context.Background(), s, actId)
	var info null.String
	err = rc.Scan(&info)
	if err != nil {
		msg.Status = -410
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	signInInfo := info.String
	if signInInfo != "" {
		//查看用户是否已经签到，如果已签到则不重复录入
		subStr := fmt.Sprintf(`{"name": "%s", "phone": "%s"}`, name.String, phone.String)
		isSignIn := strings.Index(signInInfo, subStr)
		fmt.Println("isSignIn:", isSignIn)
		if isSignIn != -1 {
			fmt.Println("已签到")
			msg.Data = []byte(fmt.Sprintf(`{"isSignIn":%t}`, true))
			service.Resp(w, &msg)
			return
		}
		//未签到，且签到列表不为空，录入信息
		s = fmt.Sprintf(`update t_activity_module set signin_info = signin_info || $1::jsonb where act_id=$2`)
	} else {
		s = fmt.Sprintf(`update t_activity_module set signin_info = $1::jsonb where act_id=$2`)
	}
	signInInfo = fmt.Sprintf(`[{"name":"%s","phone":"%s"}]`, name.String, phone.String)

	_, err = dbConn.Exec(context.Background(), s, signInInfo, actId)
	if err != nil {
		msg.Status = -411
		msg.Msg = err.Error()
		fmt.Println(msg.Status, msg.Msg)
		service.Resp(w, &msg)
		return
	}

	msg.Data = []byte(fmt.Sprintf(`{"name":"%s","phone":"%s"}`, name.String, phone.String))
	service.Resp(w, &msg)
}
