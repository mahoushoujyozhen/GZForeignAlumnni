package resume

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func DetailsWork(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -403 //r.Body 出错了
		msg.Msg = err.Error()
		logging.Error(err.Error() + "r.Body出错")
		service.Resp(w, &msg)
		return
	}

	if buf == nil {
		msg.Status = -500 //传递的值为空
		msg.Msg = "invalid/nil request param"
		service.Resp(w, &msg)
		logging.Error(err.Error() + "传值为空")
		return
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap) //unmarshal(数据编出),将json字符串解码到相应的数据结构

	if err != nil {
		msg.Status = -400 //buf解码数据出现了错误
		msg.Msg = err.Error()
		logging.Error(err.Error() + "前端head json数据解码错误")
		service.Resp(w, &msg)
		return
	}

	detailsId, err := strconv.ParseInt(fmt.Sprint(jsonMap["details_id"]), 10, 64)
	if err != nil {
		fmt.Println(err)
		msg.Status = -777
		msg.Msg = err.Error()
		logging.Error(err.Error() + "detialsWork()中jsonmap的detials_id转int64出错")
		service.Resp(w, &msg)
		return
	}

	var workInfo string

	works := make([]Work, 0)
	s := `select  work_info from t_work where id= $1 `

	rc := dbConn.QueryRow(context.Background(), s, detailsId)

	err = rc.Scan(&workInfo)

	if err != nil {
		msg.Status = -1003
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		logging.Error(err.Error() + "数据库读取错误")
		return
	}
	var data = []byte(workInfo)
	err = json.Unmarshal(data, &jsonMap)
	if err != nil {
		msg.Status = -777
		msg.Msg = err.Error()
		logging.Error(err.Error() + "detialsWork()中的work_info的json表解码错误")
		service.Resp(w, &msg)
		return
	}

	worksT := Work{
		Id:                 0,
		Position:           "",
		PositionTitle:      "",
		Organization:       "",
		OrganizationNature: "",
		OrgPhone:           "",
		Professional:       "",
		EndTime:            "",
		BeginTime:          "",
		Region:             "",
	}
	worksT.Id = detailsId
	worksT.PositionTitle = fmt.Sprint(jsonMap["position_title"])
	worksT.OrganizationNature = fmt.Sprint(jsonMap["organization_nature"])
	worksT.OrgPhone = fmt.Sprint(jsonMap["or_telephone"])
	worksT.Position = fmt.Sprint(jsonMap["position"])
	worksT.Position = fmt.Sprint(jsonMap["position"])
	worksT.Region = fmt.Sprint(jsonMap["region"])
	worksT.BeginTime = fmt.Sprint(jsonMap["begin_time"])
	worksT.EndTime = fmt.Sprint(jsonMap["end_time"])
	worksT.Organization = fmt.Sprint(jsonMap["organization"])
	worksT.Professional = fmt.Sprint(jsonMap["professional"])
	works = append(works, worksT)

	workR := WorkR{
		Status: 0,
		Data:   works,
	}

	buf, err = json.Marshal(&workR)
	if err != nil {
		msg.Status = -300
		msg.Msg = err.Error()
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, err = w.Write(buf)
	if err != nil {
		logging.Error(err.Error() + "detialsWork()数据写回错误")
	}
}
