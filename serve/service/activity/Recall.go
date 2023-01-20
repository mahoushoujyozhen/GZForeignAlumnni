package activity

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func CreateRecall(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 200,
		Msg:    "success",
	}

	dec := json.NewDecoder(r.Body)
	activity := Activity{}
	err := dec.Decode(&activity)
	if err != nil {
		msg.Status = 300
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		return
	}

	s := `update t_activity_module set act_text = $1, recall_img = $2 WHERE act_id = $3`
	_, err = dbConn.Exec(context.Background(), s, activity.Act_text, activity.Recall_img, activity.Id)
	if err != nil {
		msg.Status = 303
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		fmt.Println("Exec err=", err)
		return
	}

	service.Resp(w, &msg)
}

func DeleteRecall(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 200,
		Msg:    "success",
	}

	dec := json.NewDecoder(r.Body)
	activity := Activity{}
	err := dec.Decode(&activity)
	if err != nil {
		msg.Status = 300
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		return
	}

	s := `update t_activity_module set act_text = $1, recall_img=$2 WHERE act_id = $3`
	_, err = dbConn.Exec(context.Background(), s, activity.Act_text, activity.Recall_img, activity.Id)
	if err != nil {
		msg.Status = 303
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		fmt.Println("Exec err=", err)
		return
	}

	service.Resp(w, &msg)
}
