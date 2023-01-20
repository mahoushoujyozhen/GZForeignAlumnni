package resume

import (
	"backend/logging"
	"backend/service"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func ListMsg(w http.ResponseWriter, r *http.Request) {

	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}

	//buf, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	msg.Status = -403 //r.Body 出错了
	//	msg.Msg = err.Error()
	//	logging.Error(err.Error())
	//	service.Resp(w, &msg)
	//	return
	//}
	//
	//if buf == nil {
	//	msg.Status = -500 //传递的值为空
	//	msg.Msg = "invalid/nil request param"
	//	logging.Error(err.Error())
	//	service.Resp(w, &msg)
	//	return
	//}
	jsonMap := make(map[string]interface{})
	//err = json.Unmarshal(buf, &jsonMap) //unmarshal(数据编出),将json字符串解码到相应的数据结构
	//if err != nil {
	//	msg.Status = -400 //buf解码数据出现了错误
	//	msg.Msg = err.Error()
	//	logging.Error(err.Error())
	//	service.Resp(w, &msg)
	//	return
	//}
	//id := jsonMap["id"]

	//if id == nil {
	//	msg.Status = -500
	//	msg.Msg = "empty id ,please check again"
	//	logging.Error(err.Error())
	//	service.Resp(w, &msg)
	//	return
	//}
	var err error
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var id = session.Values["userID"].(int)

	var workInfo string
	var detailsId int64
	works := make([]Work, 0)
	s := `select  work_info,id from t_work where created_by= $1 `

	rc, err := dbConn.Query(context.Background(), s, id)

	if err != nil {
		msg.Status = -778
		msg.Msg = "t_work表的work_info查询失败"
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	for rc.Next() {

		err = rc.Scan(&workInfo, &detailsId)

		if err != nil {
			msg.Status = -1003
			msg.Msg = err.Error()
			logging.Error(err.Error())
			service.Resp(w, &msg)
			return
		}
		var data = []byte(workInfo)
		err = json.Unmarshal(data, &jsonMap)

		if err != nil {
			msg.Status = -777
			msg.Msg = err.Error()
			logging.Error(err.Error() + "listMsg()中的work_info的json表解码错误")
			service.Resp(w, &msg)
			return
		}

		worksT := Work{
			Id:           0,
			Position:     "",
			Organization: "",
			Professional: "",
			EndTime:      "",
			BeginTime:    "",
			Region:       "",
		}
		worksT.Id = detailsId
		worksT.Position = fmt.Sprint(jsonMap["position"])
		worksT.Region = fmt.Sprint(jsonMap["region"])
		worksT.BeginTime = fmt.Sprint(jsonMap["begin_time"])
		worksT.EndTime = fmt.Sprint(jsonMap["end_time"])
		worksT.Organization = fmt.Sprint(jsonMap["organization"])
		worksT.Professional = fmt.Sprint(jsonMap["professional"])
		works = append(works, worksT)

	}
	workR := WorkR{
		Status: 0,
		Data:   works,
	}

	buf, err := json.Marshal(&workR)
	if err != nil {
		msg.Status = -300
		msg.Msg = err.Error()
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	var educationInfo string

	Edus := make([]Edu, 0)
	s = `select  education_info,id from t_edu where created_by= $1 `

	rc, err = dbConn.Query(context.Background(), s, id)

	if err != nil {
		msg.Status = -778
		msg.Msg = "t_edu表的education_info查询失败"
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	for rc.Next() {

		err = rc.Scan(&educationInfo, &detailsId)

		if err != nil {
			msg.Status = -1003
			msg.Msg = err.Error()
			logging.Error(err.Error())
			service.Resp(w, &msg)
			return
		}
		var data = []byte(educationInfo)
		err = json.Unmarshal(data, &jsonMap)

		if err != nil {
			msg.Status = -777
			msg.Msg = err.Error()
			logging.Error(err.Error() + "listMsg()中的education_info的json表解码错误")
			service.Resp(w, &msg)
			return
		}

		EdusT := Edu{
			Id:                  0,
			College:             "",
			Major:               "",
			Degree:              "",
			Region:              "",
			SupportingMaterials: "",
			EndTime:             "",
			BeginTime:           "",
		}
		EdusT.Id = detailsId
		EdusT.College = fmt.Sprint(jsonMap["college"])
		EdusT.Major = fmt.Sprint(jsonMap["major"])
		EdusT.BeginTime = fmt.Sprint(jsonMap["begin_time"])
		EdusT.EndTime = fmt.Sprint(jsonMap["end_time"])
		EdusT.SupportingMaterials = fmt.Sprint(jsonMap["supporting_materials"])
		EdusT.Region = fmt.Sprint(jsonMap["region"])
		EdusT.Degree = fmt.Sprint(jsonMap["degree"])
		Edus = append(Edus, EdusT)

	}
	eduR := EduR{
		Status: 0,
		Data:   Edus,
	}
	buf, err = json.Marshal(&eduR)
	if err != nil {
		msg.Status = -300
		msg.Msg = err.Error()
		fmt.Println(err.Error())
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	//两个结构体捆绑
	total := Total{
		Status:   0,
		WorkData: workR,
		EduData:  eduR,
	}
	//设置返回头的消息类型，
	buf, err = json.Marshal(&total)
	if err != nil {
		msg.Status = -300
		msg.Msg = err.Error()
		fmt.Println(err.Error())
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")

	_, err = w.Write(buf)
	if err != nil {
		logging.Error(err.Error() + "listMsg()数据写回错误")
	}
}
