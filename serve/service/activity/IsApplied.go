package activity

import (
	"backend/service"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func IsApplied(w http.ResponseWriter, r *http.Request) {
	var err error
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	//
	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
	w.Header().Set("content-type", "application/json")             //返回数据格式是json

	r.ParseForm()
	act_id := r.FormValue("act_id")
	//user_id := r.FormValue("user_id")
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var memberId = session.Values["userID"].(int)
	aId, _ := strconv.Atoi(act_id)
	fmt.Println("aId=", act_id, "uId=", memberId)

	var userId, actId int
	isapply := false

	s := `select member_id, act_id from t_join where member_id=$1 and act_id=$2`
	rc := dbConn.QueryRow(context.Background(), s, memberId, aId)

	err = rc.Scan(&userId, &actId)
	fmt.Println("userId=", userId, "actId=", actId)
	fmt.Println(err)
	if err != nil || userId == 0 || actId == 0 {
		fmt.Println("用户未报名活动")
	} else {
		isapply = true
	}
	fmt.Println("IsApplied=", isapply)
	msg.Data = []byte(fmt.Sprintf(`
         {"isApply":%t}`, isapply))
	enc := json.NewEncoder(w)
	_ = enc.Encode(msg)

	fmt.Println("IsApplied function is successfully running")
}
