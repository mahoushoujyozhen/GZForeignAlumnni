package file

import (
	"backend/logging"
	"backend/service"
	"encoding/json"
	"fmt"
	"github.com/xuri/excelize/v2"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"
)

func ReceiveExcel(w http.ResponseWriter, r *http.Request) {
	var err error
	msg := service.ReplyProto{
		Status: 0,
		Msg:    "success",
	}
	//1.从请求头中获取文件,保存到本地
	uploadFile, fileHeader, err := r.FormFile("file")
	if err != nil {
		msg.Status = -4000
		msg.Msg = err.Error()
		logging.Error(err)
		service.Resp(w, &msg)
		return
	}

	//获取文件名
	uploadFileName := fileHeader.Filename
	//获取文件后缀
	extensionName := path.Ext(uploadFileName)

	//保存文件的路径
	var rootPath string
	rootPath, err = os.Getwd()
	filePath := fmt.Sprintf("%v\\%v", rootPath, uploadFileName)

	//保存文件到本地
	var localFile *os.File
	localFile, err = os.Create(filePath)
	defer func(localFile *os.File) {
		err := localFile.Close()
		if err != nil {

		}
	}(localFile)

	if err != nil {
		msg.Status = -4002
		msg.Msg = err.Error()
		logging.Error(err)
		service.Resp(w, &msg)
		return
	}

	_, err = io.Copy(localFile, uploadFile)
	if err != nil {
		msg.Status = -4004
		msg.Msg = err.Error()
		logging.Error(err)
		service.Resp(w, &msg)
		return
	}

	//2.获取文件名，判断是否为excel文件
	//若为excel文件则导入到数据库，否则报错

	if extensionName != ".xlsx" {
		msg.Status = -4001
		msg.Msg = "文件格式错误，仅支持xlsx文件格式"
		service.Resp(w, &msg)
		logging.Error(err)
		return
	} else {
		ReadExcel(uploadFileName, &msg)
		service.Resp(w, &msg)
	}

}

func ReadExcel(fileName string, msg *service.ReplyProto) {
	f, err := excelize.OpenFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	//插入信息模板
	var Information = []string{"序号", "会员证编号", "姓名", "性别", "民族", "籍贯", "年龄", "出生年月",
		"电话", "传真", "邮箱", "学国家/地区", "工作单位", "职称", "专业", "学历（学位）", "行政级别", "政治面貌",
		"留学时间", "留学学校", "出国工作", "访问学者", "地址", "备注"}

	fmt.Println(Information)

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		logging.Error(err)
		return
	}

	var errorRow [30]int

	for i, row := range rows {
		if i == 0 {
			//验证第一行数据是否符合模板
			//fmt.Printf("第一行数据为:%v", row)
			for j, colCell := range row {
				if colCell != Information[j] {
					msg.Msg = fmt.Sprintf("第 %v列数据不符合规范，该列列名应该为%v,该列数据插入失败\n", j+1, Information[j])
					//记录不符合规范的列
					errorRow[j] = 1
					fmt.Println(errorRow)
				}
			}
		} else {
			//插入符合规范的数据
			err := InsertExcel(row, errorRow)
			if err != nil {
				fmt.Println(err)
				logging.Error(err)
				return
			}
		}
	}
}

func InsertExcel(excelData []string, errorRow [30]int) error {
	var err error

	for i, value := range errorRow {
		if value == 1 {
			fmt.Println("errorRow:", i)
			excelData[i] = ""
		}
	}

	fmt.Println("excelData:", excelData)

	var memInfo OldMan
	memInfo.Id, _ = strconv.Atoi(excelData[1])
	memInfo.Name = excelData[2]
	memInfo.Sex = excelData[3]
	memInfo.EssentialInfo.Nation = excelData[4]      //民族,essential_info
	memInfo.EssentialInfo.NativePlace = excelData[5] // 籍贯,essential_info
	memInfo.EssentialInfo.Age = excelData[6]

	tm, err := time.Parse("01/02/2006", excelData[7])
	temp := tm.Format("2006-01-02")
	memInfo.Birthday, err = time.Parse("2006-01-05", temp)

	if err != nil {
		fmt.Println("出生日期格式转换错误，该列数据插入失败！")
		logging.Error(err)
	}

	//memInfo.Birthday, err = time.ParseInLocation("2021-01-02",excelData[7],time.Local)
	//if err != nil {
	//	fmt.Println("出生日期格式转换错误，插入该列数据失败！")
	//}

	memInfo.Phone = excelData[8]
	memInfo.EssentialInfo.Fax = excelData[9] //传真 essential_info
	memInfo.Email = excelData[10]
	memInfo.EduInfo.Region = excelData[11]              //学校国家/地区  education_info
	memInfo.WorkInfo.OrganizationNature = excelData[12] //工作单位  work_info
	memInfo.WorkInfo.PositionTitle = excelData[13]      //职称
	memInfo.WorkInfo.Professional = excelData[14]       //专业
	memInfo.EssentialInfo.Education = excelData[15]     //学历
	memInfo.EssentialInfo.SocialAppointments = excelData[16]
	memInfo.EssentialInfo.Political = excelData[17] //政治面貌
	memInfo.WorkInfo.WorkTime = excelData[18]
	memInfo.EduInfo.College = excelData[19]
	memInfo.WorkInfo.WorkType = excelData[20]
	memInfo.WorkInfo.VisitPeople = excelData[21]
	memInfo.EssentialInfo.Address = excelData[22]
	memInfo.WorkInfo.Other = excelData[23]

	memInfo.Status = 1
	memInfo.CreatedAt = time.Now()
	memInfo.UpdatedAt = time.Now()

	essentialInfoJson, _ := json.Marshal(memInfo.EssentialInfo)
	supportingMaterialJson, _ := json.Marshal(memInfo.SupportingMaterials)
	educationInfoJson, _ := json.Marshal(memInfo.EduInfo)
	workInfoJson, _ := json.Marshal(memInfo.WorkInfo)

	s1 := `insert into t_member 
	(id, name, sex, phone, email, idCard_type, idCard, birthday,
	essential_info, supporting_materials, status, created_at, updated_at)
	values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	returning id` //基本信息
	rc1 := dbConn.QueryRow(context.Background(), s1, memInfo.Id, memInfo.Name, memInfo.Sex, memInfo.Phone,
		memInfo.Email, "未填", "111111", memInfo.Birthday, essentialInfoJson, supportingMaterialJson,
		memInfo.Status, memInfo.CreatedAt, memInfo.UpdatedAt)

	var mid int64
	err = rc1.Scan(&mid)
	if err != nil {
		fmt.Println("basic err--", err)
		return err
	}

	fmt.Println("mid:--------------", mid)

	s2 := `insert into 
	t_edu(education_info,created_by)
	values($1,$2)`

	_, err = dbConn.Exec(context.Background(), s2, educationInfoJson, mid)
	if err != nil {
		fmt.Println("eduErr--:", err)
		return err
	}

	s3 := `insert into 
	t_work(work_info,created_by)
	values($1,$2)` //工作经历
	_, err = dbConn.Exec(context.Background(), s3, workInfoJson, mid)
	if err != nil {
		fmt.Println("workErr--:", err)
		return err
	}

	fmt.Println("插入完成.........")
	fmt.Println()
	fmt.Println()

	return err
}
