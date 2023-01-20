package activity

import (
	"backend/logging"
	"backend/service"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllActivities(w http.ResponseWriter, r *http.Request) {
	msg := service.ReplyProto{
		Status: 200,
		Msg:    "success",
	}

	s := `select act_id,act_name,act_profile,img_url from t_activity_module order by act_id desc`
	rows, err := dbConn.Query(context.Background(), s)
	if err != nil {
		fmt.Printf("dbConn.Query failed, err:%v\n", err)
		msg.Status = 305
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

	defer rows.Close()

	var mapSlice []map[string]interface{}

	for rows.Next() {
		var act Activity
		err := rows.Scan(&act.Id, &act.Act_name, &act.Act_profile, &act.Img_url)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			msg.Status = 304
			msg.Msg = err.Error()
			logging.Error(msg.Msg)
			service.Resp(w, &msg)
			return
		}
		item := make(map[string]interface{})
		item["act_id"] = act.Id
		item["act_name"] = act.Act_name
		item["act_profile"] = act.Act_profile
		item["img_url"] = act.Img_url

		mapSlice = append(mapSlice, item)

	}

	enc := json.NewEncoder(w)
	err = enc.Encode(mapSlice)
	if err != nil {
		fmt.Printf("Encode failed, err:%v\n", err)
		msg.Status = 301
		msg.Msg = err.Error()
		logging.Error(msg.Msg)
		service.Resp(w, &msg)
		return
	}

}
