package member

import (
	"backend/logging"
	e2 "backend/public"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func Submit(w http.ResponseWriter, r *http.Request) {
	var err error
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "注册成功，已发送管理员审核",
	}
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = e2.ERROR
		msg.Msg = e2.GetMsg(e2.ERROR)
		logging.Error(msg.Msg)
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = e2.ERROR
		msg.Msg = e2.GetMsg(e2.ERROR)
		logging.Error(msg.Msg)
		return
	}

	educationInfo := jsonMap["education_info"]
	workInfo := jsonMap["work_info"]
	essentialInfo := jsonMap["essential_info"]
	supportingMaterials := jsonMap["supporting_materials"]
	openName := jsonMap["openName"]
	password := jsonMap["password"]
	name := jsonMap["name"]
	sex := jsonMap["sex"]
	phone := jsonMap["phone"]
	email := jsonMap["email"]
	idcardType := jsonMap["idcard_type"]
	idcard := jsonMap["idcard"]
	birthday := jsonMap["birth"]
	studyType := jsonMap["study_type"]
	fileId := jsonMap["fileId"]
	supportingMaterialJson, err := json.Marshal(supportingMaterials)
	if err != nil {
		fmt.Println(err)
	}

	educationInfoJson, err := json.Marshal(educationInfo)
	if err != nil {
		fmt.Println(err)
	}

	var edu_info EduInfo
	_ = json.Unmarshal(educationInfoJson, &edu_info)

	workInfoJson, err := json.Marshal(workInfo)
	if err != nil {
		fmt.Println(err)
	}
	var work_info WorkInfo
	_ = json.Unmarshal(workInfoJson, &work_info)
	essentialInfoJson, err := json.Marshal(essentialInfo)
	if err != nil {
		fmt.Println(err)
	}
	//生成主键id
	rand.Seed(time.Now().Unix())
	timeId := 0
	for i := 0; i < 100000000; i++ {
		timeId += rand.Intn(10)
	}
	currentTimeData := time.Now()

	//检测用户是否已存在
	s := `select open_id from t_member where open_id=$1`
	rc := dbConn.QueryRow(context.Background(), s, openName)
	var id int64
	err = rc.Scan(&id)
	nonexistent := err == pgx.ErrNoRows
	if !nonexistent {
		msg.Status = 2004
		msg.Msg = "用户已存在,请直接登录!!"
		service.Resp(w, &msg)
		return
	}
	var transaction pgx.Tx
	transaction, err = dbConn.Begin(context.Background())
	defer transaction.Rollback(context.Background())

	s = `insert into t_member (id,open_id, name, sex, phone, email, idCard_type, idCard, birthday, essential_info, supporting_materials, status, created_at, updated_at,password)values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,crypt($15,gen_salt('bf'))) returning id` //基本信息
	rc = dbConn.QueryRow(context.Background(), s, timeId, openName, name, sex, phone, email, idcardType, idcard, birthday, essentialInfoJson, supportingMaterialJson, 3, currentTimeData, currentTimeData, password)
	var mid int64
	err = rc.Scan(&mid)
	if err != nil {
		msg.Status = 2000
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		fmt.Println(err)
		return
	}
	s = `insert into t_edu(education_info,created_by,study_type,file_id) values($1,$2,$3,$4) returning id`
	rc = dbConn.QueryRow(context.Background(), s, educationInfoJson, mid, studyType, fileId)
	var eid int64
	err = rc.Scan(&eid)
	if err != nil {
		msg.Status = 2001
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		fmt.Println(err)
		return
	}

	s = `insert into t_work(work_info,created_by) values($1,$2) returning id`
	rc = dbConn.QueryRow(context.Background(), s, workInfoJson, mid)
	var wid int64
	err = rc.Scan(&wid)
	if err != nil {
		msg.Status = 2002
		msg.Msg = err.Error()
		service.Resp(w, &msg)
		fmt.Println(err)
		return
	}
	_ = transaction.Commit(context.Background())
	service.Resp(w, &msg)
	return
}
