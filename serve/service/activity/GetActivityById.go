package activity

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetActivityById(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 200,
		Msg:    "success",
	}

	r.ParseForm()
	id := r.FormValue("act_id")
	var act Activity

	s := `select act_name,act_profile,img_url,start_time,end_time,act_address,apply_num from t_activity_module where act_id=$1`
	rc := dbConn.QueryRow(context.Background(), s, id)
	err := rc.Scan(&act.Act_name, &act.Act_profile, &act.Img_url, &act.Start_time, &act.End_time, &act.Act_address, &act.Apply_num)
	if err != nil {
		fmt.Printf("1scan failed, err:%v\n", err)
		msg.Status = 304
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	s1 := `select act_id,act_text,recall_img from t_activity_module where act_id=$1`
	rc1 := dbConn.QueryRow(context.Background(), s1, id)
	err = rc1.Scan(&act.Id, &act.Act_text, &act.Recall_img)
	if err != nil {
		fmt.Printf("2scan failed, err:%v\n", err)
		msg.Status = 304
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(act)
	if err != nil {
		fmt.Printf("Encode failed, err:%v\n", err)
		msg.Status = 301
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

}

//func GetActivityById(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
//	w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
//	w.Header().Set("content-type", "application/json")             //返回数据格式是json
//
//	flag := r.FormValue("flag")
//	switch flag {
//	case "true":
//		r.ParseForm()
//		id := r.FormValue("act_id")
//		fmt.Println("id=", id)
//		var act Activity
//
//		s := `select act_name,act_profile,img_url,start_time,end_time,act_address,apply_num from t_activity_module where act_id=$1`
//		rc := dbConn.QueryRow(context.Background(), s, id)
//		err := rc.Scan(&act.Act_name, &act.Act_profile, &act.Img_url, &act.Start_time, &act.End_time, &act.Act_address, &act.Apply_num)
//		if err != nil {
//			fmt.Printf("1scan failed, err:%v\n", err)
//		}
//
//		s1 := `select act_id,act_text,recall_img from t_activity_module where act_id=$1`
//		rc1 := dbConn.QueryRow(context.Background(), s1, id)
//		err = rc1.Scan(&act.Id, &act.Act_text, &act.Recall_img)
//		if err != nil {
//			fmt.Printf("2scan failed, err:%v\n", err)
//		}
//
//		fmt.Printf("act=%+v\n", act)
//
//		enc := json.NewEncoder(w)
//		_ = enc.Encode(act)
//		break
//	default:
//		w.Header().Set("Access-Control-Allow-Origin", "*")             //允许访问所有域
//		w.Header().Add("Access-Control-Allow-Headers", "Content-Type") //header的类型
//		w.Header().Set("content-type", "application/json")             //返回数据格式是json
//
//		r.ParseForm()
//
//		act_id := r.FormValue("act_id")
//		user_id := r.FormValue("user_id")
//		fmt.Println("act_id=", act_id)
//		fmt.Println("user_id=", user_id)
//		var act Activity2
//
//		s := `select act_name,act_profile,img_url,start_time,end_time,act_address,apply_num from t_activity_module where act_id=$1`
//		rc := dbConn.QueryRow(context.Background(), s, act_id)
//		err := rc.Scan(&act.Act_name, &act.Act_profile, &act.Img_url, &act.Start_time, &act.End_time, &act.Act_address, &act.Apply_num)
//		if err != nil {
//			fmt.Printf("1scan failed, err:%v\n", err)
//		}
//
//		s1 := `select act_id,act_text,recall_img from t_activity_module where act_id=$1`
//		rc1 := dbConn.QueryRow(context.Background(), s1, act_id)
//		err = rc1.Scan(&act.Id, &act.Act_text, &act.Recall_img)
//		if err != nil {
//			fmt.Printf("2scan failed, err:%v\n", err)
//		}
//
//		s2 := `select member_id, act_id from t_join where member_id=$1 and act_id=$2`
//		rc2 := dbConn.QueryRow(context.Background(), s2, user_id, act_id)
//		var userId, actId int64 = -1, -1
//		err = rc2.Scan(&userId, &actId)
//		fmt.Println("userId=", userId, "actId=", actId)
//		if userId == -1 || actId == -1 {
//			fmt.Println("用户未报名活动")
//			act.IsApplied = false
//		} else {
//			act.IsApplied = true
//		}
//		fmt.Println("IsApplied=", act.IsApplied)
//		//fmt.Println("act=", act)
//
//		enc := json.NewEncoder(w)
//		_ = enc.Encode(act)
//		fmt.Printf("act=%+v\n", act)
//		fmt.Println("GetActivityById success")
//		break
//	}
//}
