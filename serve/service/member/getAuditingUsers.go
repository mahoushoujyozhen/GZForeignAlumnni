package member

import (
	"backend/logging"
	e2 "backend/public"
	"backend/service"
	"context"
	"fmt"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strings"
)

func GetAuditingUsers(w http.ResponseWriter, r *http.Request) { //函数名改成大写
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}
	s := fmt.Sprintf(`select id,name,status,created_at from t_member where status=$1 `)
	fmt.Println(s)
	rs, err := dbConn.Query(context.Background(), s, 3)
	if err != nil {
		msg.Status = e2.ERROR_NOT_EXIST_USER         //自定义的错误状态码，自己去pkg.e中去定义,尽量自己模块的状态码都定义成一起
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_USER) //自定义错误信息
		logging.Error(msg.Msg)                       //错误时输出日志
		service.Resp(w, &msg)
		fmt.Println(err)
		return
	}

	var users []string
	for rs.Next() {
		var id, status int64
		var name, date null.String //sql.NullString类型
		//将数据库查询到的值对应到相应的变量里
		err = rs.Scan(&id, &name, &status, &date)
		//如果err不为空
		if err != nil {
			msg.Status = -1003
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
			return
		}
		//将每个用户信息追加到users数组里
		users = append(users, fmt.Sprintf(`{"id":%d,"name":"%s","status":%d,"date":"%s"}`,
			id, name.String, status, date.String))
		fmt.Println(name)
	}
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(users, ",")))
	//测试
	//fmt.Println(string(msg.Data))
	//fmt.Println(users)
	service.Resp(w, &msg)
}
