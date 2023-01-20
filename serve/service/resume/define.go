package resume

import (
	"backend/service"
)

var dbConn = service.GetDatabaseConnection()

type Work struct {
	Id                 int64  `json:"id"`
	Position           string `json:"position"`
	PositionTitle      string `json:"position_title"`
	Organization       string `json:"organization"`
	OrganizationNature string `json:"organization_nature"`
	Professional       string `json:"professional"`
	BeginTime          string `json:"begin_Time"`
	EndTime            string `json:"end_Time"`
	Region             string `json:"region"`
	OrgPhone           string `json:"org_phone"`
}

type WorkR struct {
	Status int    `json:"status"`
	Data   []Work `json:"data"`
}

type Edu struct {
	FileId              string `json:"file_id"`
	Id                  int64  `json:"id"`
	StudyType           string `json:"studyType"`
	College             string `json:"college"`
	Major               string `json:"major"`
	Degree              string `json:"degree"`
	Region              string `json:"region"`
	SupportingMaterials string `json:"supporting_materials"`
	BeginTime           string `json:"begin_time"`
	EndTime             string `json:"end_time"`
}

type EduR struct {
	Status int   `json:"status"`
	Data   []Edu `json:"data"`
}

type Total struct {
	Status   int   `json:"status"`
	WorkData WorkR `json:"work_data"`
	EduData  EduR  `json:"edu_data"`
}
