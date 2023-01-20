package activity

import (
	"backend/service"
	"backend/util"
	"context"
	"fmt"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strings"
	"time"
)

//获取我的活动列表信息
//获取用户参与的所有活动信息
func GetActMsg(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	if strings.ToLower(r.Method) != "get" {
		msg.Status = -400
		msg.Msg = "invalid request,should be get"
		service.Resp(w, &msg)
		return
	}
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var memberId = session.Values["userID"].(int)
	//memberId := r.URL.Query().Get("id")
	//if memberId == "" {
	//	msg.Status = -1010
	//	msg.Msg = "missing id"
	//	service.Resp(w, &msg)
	//	return
	//}

	s := `select act_id from t_join where member_id=$1`
	rs, err := dbConn.Query(context.Background(), s, memberId) //读取每一个活动id
	if err != nil {
		msg.Status = -401
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		return

	}
	var actmsg []interface{} //用于记录每个活动的id
	var count = 0
	for rs.Next() { //将每一个活动id添加到 actmsg切片中
		var actId int
		err = rs.Scan(&actId)
		if err != nil {
			msg.Status = -402
			msg.Msg = err.Error()
			service.Resp(w, &msg)
			return
		}
		actmsg = append(actmsg, actId)
		count++
	}
	if count < 1 { //如果读取的活动id小于1，则直接返回
		msg.Status = -403
		msg.Msg = "you have no activity"
		service.Resp(w, &msg)
		return
	}
	var allactmsg []string     //用于存储每一个活动的所需信息
	for _, v := range actmsg { //通过每一个活动的 id 获取活动的详细信息
		s = `select act_name, start_time, end_time, act_address, apply_num from t_activity_module where act_id=$1`
		rc := dbConn.QueryRow(context.Background(), s, v)
		var isOver bool     //活动未结束
		var isselect = true //向前端传递可显示信号
		var applyNum int64
		var actName, actAddress, startTime, endTime null.String
		err = rc.Scan(&actName, &startTime, &endTime, &actAddress, &applyNum)
		if err != nil {
			fmt.Println("here")
			msg.Status = -404
			msg.Msg = err.Error()
			service.Resp(w, &msg)
			return
		}
		//将活动时间与现在进行比较，返回bool，为true即活动未结束；false表示活动已结束
		now := time.Now()
		loc, _ := time.LoadLocation("Local")
		theTime, err := time.ParseInLocation("2006-01-02 15:04", endTime.String, loc)
		if err == nil && now.Before(theTime) {
			isOver = true
		} else if now.After(theTime) {
			isOver = false
		}
		allactmsg = append(allactmsg, fmt.Sprintf(`{"act_id":"%d", "act_name":"%s","start_time":"%s","end_time":"%s","act_address":"%s", "apply_num":"%d", "isOver":"%t", "isselect":"%t"}`,
			v, actName.String, startTime.String, endTime.String, actAddress.String, applyNum, isOver, isselect))
	}
	//活动先后排序
	for i := 1; i < count; i++ {
		for j := 0; j < count-i; j++ {
			if actmsg[j].(int) < actmsg[j+1].(int) {
				actmsg[j], actmsg[j+1] = actmsg[j+1], actmsg[j]
				allactmsg[j], allactmsg[j+1] = allactmsg[j+1], allactmsg[j]
			}
		}
	}
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(allactmsg, ",")))
	fmt.Println("GetActMsg function is successfully running")
	service.Resp(w, &msg)
}
