package member

import (
	"backend/service"
	"github.com/jackc/pgtype"
)

var dbConn = service.GetDatabaseConnection()

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

type WorkInfo struct {
	Positon            string `json:"positon"`
	PositonTitle       string `json:"positon_title"`
	Professional       string `json:"professional"`
	Organization       string `json:"organization"`
	OrganizationNature string `json:"organization_nature"`
	OrTelephone        string `json:"or_telephone"`
	Region             string `json:"region"`
	BeginTime          string `json:"begin_time"`
	EndTime            string `json:"end_time"`
}

type EduInfo struct {
	Major     string `json:"major"`
	College   string `json:"college"`
	Region    string `json:"region"`
	BeginTime string `json:"begin_time"`
	EndTime   string `json:"end_time"`
	StudyType string `json:"study_type"`
}
