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

func DetailsEdu(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -403 //r.Body 出错了
		msg.Msg = err.Error()
		fmt.Println("DetailsEdu()中r.Body读取错误 ")
		logging.Error(err.Error() + "DetailsEdu()中r.Body读取错误 ")
		service.Resp(w, &msg)
		return
	}

	if buf == nil {
		msg.Status = -500 //传递的值为空
		msg.Msg = "invalid/nil request param 传递的值为空"
		logging.Error(err.Error() + "invalid/nil request param 传递的值为空")
		service.Resp(w, &msg)
		return
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap) //unmarshal(数据编出),将json字符串解码到相应的数据结构

	if err != nil {
		msg.Status = -400 //buf解码数据出现了错误
		msg.Msg = err.Error()
		logging.Error(err.Error() + "detailsEdu()的josn.Unmarshal解码错误")
		service.Resp(w, &msg)
		return
	}

	detailsId, err := strconv.ParseInt(fmt.Sprint(jsonMap["details_id"]), 10, 64)
	if err != nil {
		fmt.Println(err)
		msg.Status = -777
		msg.Msg = err.Error()
		logging.Error(err.Error() + "detialsEdu()中jsonmap的detialsId转int64出错")
		service.Resp(w, &msg)
		return
	}

	var eduInfo string
	var fileId string
	var studyType string
	edus := make([]Edu, 0)

	//少了文件系统会发现没有file_id,前端会报错，can't scan into dest[1]: cannot scan null into *string
	s := `select  education_info,file_id,study_type from t_edu where id= $1 `
	rc := dbConn.QueryRow(context.Background(), s, detailsId)
	err = rc.Scan(&eduInfo, &fileId, &studyType)

	//确保图片显示错误，页面其他信息正常显示
	if err != nil {
		//处理缺失文件id file_id
		s = `select  education_info,study_type from t_edu where id= $1`
		rc := dbConn.QueryRow(context.Background(), s, detailsId)
		err = rc.Scan(&eduInfo, &studyType)
		if err != nil {
			msg.Status = -1003
			msg.Msg = err.Error()
			fmt.Println(err.Error() + "detialsEdu()t_edu查找json表education_info出错")
			service.Resp(w, &msg)
			logging.Error(err.Error() + "detailsEdu()t_edu查找json表education_info出错")
			return
		}
		var data = []byte(eduInfo)
		err = json.Unmarshal(data, &jsonMap)
		if err != nil {
			msg.Status = -777
			msg.Msg = err.Error()
			logging.Error(err.Error() + "detialsEdu()中的education_info的json表解码错误")
			service.Resp(w, &msg)
			return
		}

		eduT := Edu{
			FileId:              "",
			StudyType:           "",
			Id:                  0,
			College:             "",
			Major:               "",
			Degree:              "",
			Region:              "",
			SupportingMaterials: "",
			EndTime:             "",
			BeginTime:           "",
		}
		eduT.Id = detailsId
		eduT.StudyType = studyType
		eduT.College = fmt.Sprint(jsonMap["college"])
		eduT.Major = fmt.Sprint(jsonMap["major"])
		eduT.Degree = fmt.Sprint(jsonMap["degree"])
		eduT.Region = fmt.Sprint(jsonMap["region"])
		eduT.SupportingMaterials = fmt.Sprint(jsonMap["supporting_materials"])
		eduT.BeginTime = fmt.Sprint(jsonMap["begin_time"])
		eduT.EndTime = fmt.Sprint(jsonMap["end_time"])
		edus = append(edus, eduT)

		EduR := EduR{
			Status: 0,
			Data:   edus,
		}

		buf, err = json.Marshal(&EduR)
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
			logging.Error(err.Error() + "detialsEdu()数据写回错误")
		}

		return
	}

	//处理存在文件id file_id
	var data = []byte(eduInfo)
	err = json.Unmarshal(data, &jsonMap)
	eduT := Edu{
		FileId:              "",
		Id:                  0,
		College:             "",
		Major:               "",
		Degree:              "",
		Region:              "",
		SupportingMaterials: "",
		EndTime:             "",
		BeginTime:           "",
	}
	eduT.FileId = fileId
	eduT.Id = detailsId
	eduT.StudyType = studyType

	eduT.College = fmt.Sprint(jsonMap["college"])
	eduT.Major = fmt.Sprint(jsonMap["major"])
	eduT.Degree = fmt.Sprint(jsonMap["degree"])
	eduT.Region = fmt.Sprint(jsonMap["region"])
	eduT.SupportingMaterials = fmt.Sprint(jsonMap["supporting_materials"])
	eduT.BeginTime = fmt.Sprint(jsonMap["begin_time"])
	eduT.EndTime = fmt.Sprint(jsonMap["end_time"])
	edus = append(edus, eduT)

	EduR := EduR{
		Status: 0,
		Data:   edus,
	}

	buf, err = json.Marshal(&EduR)
	if err != nil {
		msg.Status = -300
		msg.Msg = err.Error()
		logging.Error(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	_, err = w.Write(buf)
	if err != nil {
		logging.Error(err.Error() + "detialsEdu()数据写回错误")
	}
}
