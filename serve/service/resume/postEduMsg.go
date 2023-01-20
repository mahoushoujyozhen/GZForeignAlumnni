package resume

import (
	"backend/logging"
	"backend/service"
	"backend/util"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostEduMsg(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -403 //r.Body 出错了
		msg.Msg = err.Error()
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	if buf == nil {
		msg.Status = -500 //传递的值为空
		msg.Msg = "invalid/nil request param"
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}
	jsonMap := make(map[string]interface{})
	//unmarshal(数据编出),将json字符串解码到相应的数据结构
	err = json.Unmarshal(buf, &jsonMap)

	if err != nil {
		msg.Status = -400 //buf解码数据出现了错误
		msg.Msg = err.Error()
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	//id := jsonMap["id"]
	store := util.GetStore()
	session, _ := store.Get(r, "sessionID")
	var id = session.Values["userID"].(int)

	fileId := jsonMap["file_id"]

	//删去jsonMap中的id
	//delete(jsonMap, "id")
	delete(jsonMap, "file_id")

	//获取study_type
	studyType := jsonMap["studyType"]

	delete(jsonMap, "studyType")

	//转json
	buf, err = json.Marshal(jsonMap)

	if err != nil {
		msg.Status = -300
		msg.Msg = err.Error()
		fmt.Println(err.Error())
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}

	//显示json的string形式
	//fmt.Println(string(buf))

	//开始插入buf
	s := `insert into t_edu(education_info,created_by,file_id,study_type) values($1,$2,$3,$4)`

	_, err = dbConn.Exec(context.Background(), s, buf, id, fileId, studyType)
	if err != nil {
		msg.Status = 550
		msg.Msg = err.Error()
		logging.Error(err.Error())
		service.Resp(w, &msg)
		return
	}
	msg.Msg = "work_info insert success"
	service.Resp(w, &msg)
	return
}
