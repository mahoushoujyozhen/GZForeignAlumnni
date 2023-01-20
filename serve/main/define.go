package main

import (
	"backend/service/activity"

	"backend/service/file"
	"backend/service/member"
	"backend/service/memberCenter"
	"backend/service/message"
	"backend/service/resume"
)

var routeRegx string = "^/(?P<type>\\w+)/(?P<target>\\w+)/(?P<action>\\w+)(\\?(?P<query>.*))?" // URI正则表达式

var routeMap = map[string]interface{}{

	"/user/get":           member.GetAuditingUsers,
	"/admin/listuser":     message.ListUser,
	"/admin/searchuser":   message.SearchUser,
	"/admin/listalluser":  message.ListAllUser,
	"/admin/listreceiver": message.ListReceiver,
	"/message/ws":         message.MessageService,

	"/file/download": file.DownloadService,

	"/user/listmsg":     resume.ListMsg,
	"/user/detailswork": resume.DetailsWork,
	"/user/detailsedu":  resume.DetailsEdu,
	"/user/postworkmsg": resume.PostWorkMsg,
	"/user/postedumsg":  resume.PostEduMsg,
	"/user/showmsg":     memberCenter.Showmsg,
	"/user/changemsg":   memberCenter.Changemsg,

	"/admin/act_release":  activity.Act_release,
	"/admin/act_delete":   activity.Act_delete,
	"/admin/createRecall": activity.CreateRecall,
	"/admin/deleteRecall": activity.DeleteRecall,
	"/admin/act_change":   activity.Act_change,
	"/activity/getAll":    activity.GetAllActivities,
	"/activity/getById":   activity.GetActivityById,
	"/admin/signinList":   activity.SigninList,
	"/admin/getQrcode":    activity.GetQrcode,
	"/user/signIn":        activity.SignIn,
	"/user/apply":         activity.ApplyActivity,
	"/user/cancel":        activity.CancelActivity,
	"/user/isapply":       activity.IsApplied,
	"/user/peractmsg":     activity.GetActMsg,

	"/user/nouser":     member.GetUserEd,
	"/user/already":    member.Audited,
	"/user/userchange": member.ChangeUsers,

	"/user/logout": member.Logout,
	"/user/submit": member.Submit,

	"/admin/getUser":       member.GetAuditingUsers,
	"/admin/userInfo":      member.GetUserInfo,
	"/admin/opinionSubmit": member.SubmitAuditOpinion,

	"/admin/import_excel": file.ReceiveExcel,
	"/user/getopinion":    memberCenter.GetOpinion,
}
