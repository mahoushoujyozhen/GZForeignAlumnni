package activity

import (
	"backend/service"
	"gopkg.in/guregu/null.v4"
)

var dbConn = service.GetDatabaseConnection()

//活动结构体

type Activity struct {
	Id          int64       `json:"id"`
	Act_text    null.String `json:"act_text"`
	Act_name    null.String `json:"act_name"`
	Act_profile null.String `json:"act_profile"`
	Img_url     null.String `json:"img_url"`
	Start_time  null.String `json:"start_time"`
	End_time    null.String `json:"end_time"`
	Act_address null.String `json:"act_address"`
	Apply_num   int64       `json:"apply_num"`
	IsApplied   bool        `json:"isApplied"`
	Recall_img  null.String `json:"recall_img"`
}

//获取签到列表信息
type MsgRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Qrcode 活动二维码结构体
type Qrcode struct {
	Status int `json:"status"`

	Msg string `json:"msg,omitempty"`

	Qrcode []byte `json:"qrcode,omitempty"`
}
