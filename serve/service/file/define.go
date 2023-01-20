package file

import (
	"backend/logging"
	"backend/public"
	"backend/service"
	"backend/util"
	"github.com/jackc/pgtype"
	"time"
)

var rootPath = "/resources"

//var rootPath = "D:\\Git\\intergrations\\alumni-association\\serve"
var Snowflake *util.Snowflake

func init() {
	var errors error
	// 创建唯一Id生成器
	Snowflake, errors = util.GetSnowflake(time.Now().UnixMilli(), 1, 1, 0)
	// 处理创建失败的情况
	if errors != nil {
		logging.Error(public.ERR_SNOWFLAKE_CREATE)
	} else {
		logging.Info(public.MSG_SNOWFLAKE_CREATE)
	}
}

var dbConn = service.GetDatabaseConnection()

type OldMan struct {
	Id                  int                 `json:"id"`
	Phone               string              `json:"phone"`
	Name                string              `json:"name"`
	Sex                 string              `json:"sex"`
	Email               string              `json:"email"`
	IdcardType          string              `json:"idcard_type"`
	Idcard              string              `json:"idcard"`
	Birthday            time.Time           `json:"birthday"`
	EssentialInfo       EssentialInfo       `json:"essential_info"`
	SupportingMaterials SupportingMaterials `json:"supporting_materials"`
	Status              int                 `json:"status"`
	CreatedAt           time.Time           `json:"created_at"`
	UpdatedAt           time.Time           `json:"updated_at"`
	WorkId              int                 `json:"work_id"`
	WorkInfo            WorkInfo            `json:"work_info"`
	EduId               int                 `json:"edu_id"`
	EduInfo             EduInfo             `json:"education_info"`
	StudyType           int                 `json:"study_type"`
}

type EssentialInfo struct {
	Nation             string `json:"nation"`       // 民族
	NativePlace        string `json:"native_place"` //籍贯
	Nationality        string `json:"nationality"`  //国籍
	Age                string `json:"age"`
	Education          string `json:"education"` //学历
	Region             string `json:"region"`    //
	Address            string `json:"address"`
	Fax                string `json:"fax"`
	Political          string `json:"political"`
	SocialAppointments string `json:"social_appointments"`
}

type SupportingMaterials struct {
	Relation string `json:"relation"`
	Content  string `json:"content"`
}

type WorkInfo struct {
	Position           string `json:"position"`       //职位
	PositionTitle      string `json:"position_title"` //职称
	Professional       string `json:"professional"`   //专业
	Region             string `json:"region"`
	OrganizationNature string `json:"organization_nature"`
	WorkTime           string `json:"work_time"`
	VisitPeople        string `json:"visit_people"`
	WorkType           string `json:"work_type"`
	Other              string `json:"other"`
}

type EduInfo struct {
	Major     string    `json:"major"`
	College   string    `json:"college"`
	Region    string    `json:"region"`
	BeginTime time.Time `json:"begin_time"`
	EndTime   time.Time `json:"end_time"`
}

type MemInfo struct {
	Id                  int         `json:"id"`
	Phone               string      `json:"phone"`
	Name                string      `json:"name"`
	Sex                 string      `json:"sex"`
	Email               string      `json:"email"`
	IdcardType          string      `json:"idcard_type"`
	Idcard              string      `json:"idcard"`
	Birthday            pgtype.Date `json:"birthday"`
	EssentialInfo       pgtype.JSON `json:"essential_info"`
	SupportingMaterials pgtype.JSON `json:"supporting_materials"`
	Status              int         `json:"status"`
	CreatedAt           pgtype.Date `json:"created_at"`
	UpdatedAt           pgtype.Date `json:"updated_at"`
	WorkId              int         `json:"work_id"`
	WorkInfo            pgtype.JSON `json:"work_info"`
	EduId               int         `json:"edu_id"`
	EducationInfo       pgtype.JSON `json:"education_info"`
	StudyType           int         `json:"study_type"`
}
