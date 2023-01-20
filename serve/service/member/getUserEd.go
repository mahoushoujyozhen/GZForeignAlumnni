package member

import (
	"backend/logging"
	"backend/service"
	"context"
	"fmt"
	"gopkg.in/guregu/null.v4"
	"net/http"
	"strings"
)

func GetUserEd(w http.ResponseWriter, r *http.Request) {
	var err error
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	} //返回一个标志，不让前端报错
	var id int64
	var name string
	var status int
	var created_at null.String

	s := `select id,name,status,created_at from t_member order by id`
	rs, err := dbConn.Query(context.Background(), s)
	if err != nil {
		fmt.Println(err)
		msg.Status = -1000
		msg.Msg = err.Error()
		logging.Error(msg.Msg) //错误时输出日志
		service.Resp(w, &msg)
		return
	}
	var users []string
	for rs.Next() {
		err = rs.Scan(&id, &name, &status, &created_at)
		if err != nil {
			fmt.Println(err)
			msg.Status = -1003
			msg.Msg = err.Error()
			logging.Error(msg.Msg) //错误时输出日志
			service.Resp(w, &msg)
			return
		}
		if status == 3 {
			continue
		}
		users = append(users, fmt.Sprintf(`{"id":"%d","name":"%s","status":"%d","created_at":"%s"}`, id, name, status, created_at.String))
	}
	msg.Data = []byte(fmt.Sprintf(`[%s]`, strings.Join(users, ",")))
	fmt.Println("传输完毕")
	service.Resp(w, &msg)
	return

}
