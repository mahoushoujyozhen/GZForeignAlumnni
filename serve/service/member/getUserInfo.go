package member

import (
	e2 "backend/public"
	"backend/service"
	"context"
	"fmt"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strconv"
)

// GetUserInfo 获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}
	//获取url中的id的值
	id := r.URL.Query().Get("id")
	fmt.Println("后端拿到的id---------：", id)
	if id == "" {
		msg.Status = e2.ERROR_NOT_EXIST_USER_ID
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_USER_ID)
		service.Resp(w, &msg)
		return
	}
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_STATUS)
		msg.Status = e2.ERROR_NOT_EXIST_STATUS
		service.Resp(w, &msg)
		return
	}

	//sql语句
	s := fmt.Sprintf(`select phone,sex,email,idCard,birthday,essential_info,supporting_materials,work_info,education_info,file_id ,study_type,idcard_type from v_member where id=%d`, uid)
	//执行sql语句
	//fmt.Println("正在查询" + s)
	rs, err := dbConn.Query(context.Background(), s)
	if err != nil {
		msg.Status = e2.ERROR_NOT_EXIST_USER_INFO
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_USER_INFO)
		service.Resp(w, &msg)
		return
	}
	var phone, sex, email, idCard, birthday, essential_info, supportial_material, work_info, education_info, file_id, study_type, idcard_type null.String //sql.NullString类型
	for rs.Next() {
		//将数据库查询到的值对应到相应的变量里
		err = rs.Scan(&phone, &sex, &email, &idCard, &birthday, &essential_info, &supportial_material, &work_info, &education_info, &file_id, &study_type, &idcard_type)
		//如果err不为空
		if err != nil {
			fmt.Println(err)
			msg.Status = e2.ERROR_DATABASE_SEARCH_FAIL
			msg.Msg = e2.GetMsg(e2.ERROR_DATABASE_SEARCH_FAIL)
			service.Resp(w, &msg)
			return
		}
		//将mName转化为json形式
		userinfo := fmt.Sprintf(`{"phone":"%s","sex":"%s","email":"%s","idCard":"%s","birthday":"%s","essential_info":%s,"supportial_material":%s,"work_info":%s,"education_info":%s,"file_id":"%s","study_type":"%s" ,"idcard_type":"%s"}`,
			phone.String, sex.String, email.String, idCard.String, birthday.String, essential_info.String, supportial_material.String, work_info.String, education_info.String, file_id.String, study_type.String, idcard_type.String)
		if err != nil {
			fmt.Println("MapToJsonDemo err", err)
		}
		fmt.Println(string(userinfo))
		msg.Data = []byte(string(userinfo))
	}
	service.Resp(w, &msg)
}
