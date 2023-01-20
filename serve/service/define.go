package service

import (
	"github.com/jmoiron/sqlx/types"
)

type ReplyProto struct {
	//Status, 0: success, others: fault
	Status int `json:"status"`

	//Msg, Action result describe by literal
	Msg string `json:"msg,omitempty"`

	//Data, operand
	Data types.JSONText `json:"data,omitempty"`

	//AllUser, storage the allUserList
	AllUser types.JSONText `json:"alluser,omitempty"`

	//PageCount ,just page count
	PageCount int64 `json:"pageCount,omitempty"`

	// RowCount, just row count
	RowCount int64 `json:"rowCount,omitempty"`

	//API, call target
	API string `json:"API,omitempty"`

	//Method, using http method
	Method string `json:"method,omitempty"`

	//SN, call order
	SN int `json:"SN,omitempty"`

	Qrcode []byte `json:"qrcode,omitempty"`
}

type Member_msg struct {
	Status     int    `json:"status"`
	Msg        string `json:"msg"`
	Name       string `json:"username"`
	Id         int    `json:"userid"`
	Phone      string `json:"userphone"`
	Email      string `json:"useremail"`
	UserStatus int    `json:"user_status"`
	Sex        string `json:"usersex"`
	Nickname   string `json:"nickname"`
	IdCard     string `json:"id_card"`
	IdCardType string `json:"id_card_type"`
}

type NotificationBean struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title,omitempty"`
	Time    int64  `json:"time,omitempty"`
	Content string `json:"content,omitempty"`
}

type ResponseBean struct {
	StatusCode int         `json:"statusCode,omitempty"`
	Message    string      `json:"message,omitempty"`
	Details    interface{} `json:"details,omitempty"`
}

func ResponseBeanFactory(statusCode int, message string, details interface{}) *ResponseBean {
	return &ResponseBean{StatusCode: statusCode, Message: message, Details: details}
}
