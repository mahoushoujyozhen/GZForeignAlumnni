package member

import (
	e2 "backend/public"
	"backend/service"
	"context"
	"fmt"
	"net/http"
	"strconv"
)

// SubmitAuditOpinion 审核通过/未通过更改用户标志位
func SubmitAuditOpinion(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: http.StatusOK,
		Msg:    "success",
	}

	id := r.URL.Query().Get("id")
	if id == "" {
		msg.Status = e2.ERROR_NOT_EXIST_USER_ID
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_USER_ID)
		service.Resp(w, &msg)
		return
	}
	uid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		msg.Msg = e2.GetMsg(e2.ERROR_USER_ID)
		msg.Status = e2.ERROR_USER_ID
		service.Resp(w, &msg)
		return
	}
	var values []interface{}

	values = append(values, uid)

	opinion := r.URL.Query().Get("opinion")
	if opinion == "" {
		msg.Status = e2.ERROR_NOT_EXIST_OPINION
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_OPINION)
		service.Resp(w, &msg)
		return
	}
	values = append(values, opinion)
	status := r.URL.Query().Get("status")
	if status == "" {
		msg.Status = e2.ERROR_NOT_EXIST_OPINION
		msg.Msg = e2.GetMsg(e2.ERROR_NOT_EXIST_OPINION)
		service.Resp(w, &msg)
		return
	}
	ustatus, err := strconv.ParseInt(status, 10, 64)
	if err != nil {
		msg.Status = e2.ERROR_STATUS
		msg.Msg = e2.GetMsg(e2.ERROR_STATUS)
		service.Resp(w, &msg)
		return
	}

	values = append(values, ustatus)
	s := fmt.Sprintf(`update t_member set opinion=$2,status=$3 where id=$1`)
	fmt.Println(s)
	_, err = dbConn.Exec(context.Background(), s, values...)
	if err != nil {
		msg.Status = e2.ERROR_DATABASE_UPDATE_FAIL
		msg.Msg = e2.GetMsg(e2.ERROR_DATABASE_UPDATE_FAIL)
		service.Resp(w, &msg)
		return
	}
	msg.Data = []byte(fmt.Sprintf(`{"id":%d}`, uid))
	service.Resp(w, &msg)
}
