package member

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v4"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func Audited(w http.ResponseWriter, r *http.Request) {
	var err error

	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	} //返回一个标志，不让前端报错
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		msg.Status = -403
		msg.Msg = err.Error()
		return
	}

	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(buf, &jsonMap)
	if err != nil {
		msg.Status = -400
		msg.Msg = err.Error()
		return
	}
	id := jsonMap["id"]
	fmt.Println("shenhewei")
	var open_id string
	var name string
	var sex string
	var phone string
	var email string
	var idcard_type string
	var idcard string
	var birthday time.Time
	var essential_info string
	var supporting_materials string
	var status int
	var created_at time.Time
	var updated_at time.Time
	var work_info string
	var created_by int
	var education_info string
	var study_type string
	var createds_by int
	var uuu int
	var opinion string
	s := `select id,open_id,name,sex,phone,email,idcard_type,
    idcard,birthday,essential_info,supporting_materials,status,created_at,
    updated_at,opinion from t_member where id=$1`
	rc := dbConn.QueryRow(context.Background(), s, id)
	err = rc.Scan(&id, &open_id, &name, &sex, &phone, &email, &idcard_type, &idcard, &birthday, &essential_info, &supporting_materials,
		&status, &created_at, &updated_at, &opinion)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
		msg.Status = -1000
		msg.Msg = err.Error()
		logging.Error(msg.Msg) //错误时输出日志
		service.Resp(w, &msg)
		return
	}
	fmt.Println(birthday)
	opinion = strings.Replace(opinion, "\n", "\\n", -1)
	essential_info = strings.Replace(essential_info, "{", "", -1)
	essential_info = strings.Replace(essential_info, "}", "", -1)
	supporting_materials = strings.Replace(supporting_materials, "{}", "\"\"", -1)
	supporting_materials = strings.Replace(supporting_materials, "[{", "[x", -1)
	supporting_materials = strings.Replace(supporting_materials, "}]", "x]", -1)
	supporting_materials = strings.Replace(supporting_materials, "{", "", -1)
	supporting_materials = strings.Replace(supporting_materials, "}", "", -1)
	supporting_materials = strings.Replace(supporting_materials, "[x", "[{", -1)
	supporting_materials = strings.Replace(supporting_materials, "x]", "}]", -1)

	s1 := `select work_info,created_by from t_work where created_by=$1`
	rc1 := dbConn.QueryRow(context.Background(), s1, id)
	err = rc1.Scan(&work_info, &created_by)
	nonexistent := err == pgx.ErrNoRows
	if nonexistent {
		fmt.Println("没找到！")
		fmt.Println(err)
		uuu = 1

	}
	work_info = strings.Replace(work_info, "{", "", -1)
	work_info = strings.Replace(work_info, "}", "", -1)
	work_info = strings.Replace(work_info, "region", "w_region", -1)
	fmt.Println(work_info)
	/*var education_info string
	  var study_type int
	  var createds_by int*/
	s2 := `select education_info,study_type,created_by from t_edu where created_by=$1`
	rc2 := dbConn.QueryRow(context.Background(), s2, id)
	err = rc2.Scan(&education_info, &study_type, &createds_by)
	nonexistent1 := err == pgx.ErrNoRows
	if nonexistent1 {
		fmt.Println("没找到！")
		fmt.Println("3")
		fmt.Println(err)
		if uuu == 1 {
			uuu = 11
		} else {
			uuu = 2
		}
	}
	education_info = strings.Replace(education_info, "{}", "\"\"", -1)
	education_info = strings.Replace(education_info, "[{", "[[x", -1)
	education_info = strings.Replace(education_info, "}]", "x]]", -1)
	education_info = strings.Replace(education_info, "{", "", -1)
	education_info = strings.Replace(education_info, "}", "", -1)
	education_info = strings.Replace(education_info, "[[x", "[{", -1)
	education_info = strings.Replace(education_info, "x]]", "}]", -1)

	var users []string
	if uuu == 0 {
		users = append(users, fmt.Sprintf(`{"id":"%d","open_id":"%s","name":"%s","sex":"%s","phone":"%s",
"email":"%s","idcard_type":"%s","idcard":"%s","birthday":"%s",%s,%s,"status":"%d","created_at":"%s","updated_at":"%s","opinion":"%s",%s,
"created_by":"%d",%s,"study_type":"%s","createds_by":"%d"}`,
			id, open_id, name, sex, phone, email, idcard_type, idcard, birthday, essential_info, supporting_materials,
			status, created_at, updated_at, opinion, work_info, created_by, education_info, study_type, createds_by))
	} else if uuu == 1 {
		users = append(users, fmt.Sprintf(`{"id":"%d","open_id":"%s","name":"%s","sex":"%s","phone":"%s",
"email":"%s","idcard_type":"%s","idcard":"%s","birthday":"%s",%s,%s,"status":"%d","created_at":"%s","updated_at":"%s","opinion":"%s",%s,"study_type":"%s","createds_by":"%d"}`,
			id, open_id, name, sex, phone, email, idcard_type, idcard, birthday, essential_info, supporting_materials,
			status, created_at, updated_at, opinion, education_info, study_type, createds_by))
	} else if uuu == 2 {
		users = append(users, fmt.Sprintf(`{"id":"%d","open_id":"%s","name":"%s","sex":"%s","phone":"%s",
"email":"%s","idcard_type":"%s","idcard":"%s","birthday":"%s",%s,%s,"status":"%d","created_at":"%s","updated_at":"%s","opinion":"%s",%s,
"created_by":"%d"}`,
			id, open_id, name, sex, phone, email, idcard_type, idcard, birthday, essential_info, supporting_materials,
			status, created_at, updated_at, opinion, work_info, created_by))

	} else {
		users = append(users, fmt.Sprintf(`{"id":"%d","open_id":"%s","name":"%s","sex":"%s","phone":"%s",
"email":"%s","idcard_type":"%s","idcard":"%s","birthday":"%s",%s,%s,"status":"%d","created_at":"%s","updated_at":"%s","opinion":"%s"}`,
			id, open_id, name, sex, phone, email, idcard_type, idcard, birthday, essential_info, supporting_materials,
			status, created_at, updated_at, opinion))
	}
	fmt.Println(users)

	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(users, ",")))
	fmt.Println("传输完毕s1")
	service.Resp(w, &msg)
	return

}
