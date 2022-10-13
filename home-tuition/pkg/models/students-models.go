package models

import (
	"github.com/jinzhu/gorm"
	"github.com/rebontadeb/home-tuition/pkg/config"
)

var db *gorm.DB

type Students struct {
	//gorm.Model
	StudentId   int    `gorm:"column:studentId"`
	FirstName   string `gorm:"column:FirstName"`
	LastName    string `gorm:"column:LastName"`
	Address     string `gorm:"column:Address"`
	City        string `gorm:"column:City"`
	PhoneNumber string `gorm:"column:phoneNumber"`
	SchoolName  string `gorm:"column:schoolName"`
	Class       string `gorm:"column:class"`
}

type NewStudents struct {
	FirstName   string `gorm:"column:FirstName"`
	LastName    string `gorm:"column:LastName"`
	Address     string `gorm:"column:Address"`
	City        string `gorm:"column:City"`
	PhoneNumber string `gorm:"column:phoneNumber"`
	SchoolName  string `gorm:"column:schoolName"`
	Class       string `gorm:"column:class"`
}

func init() {
	config.DbConnection()
	db = config.GetDbConnection()
	db.AutoMigrate(&Students{})
}

func (s *NewStudents) CreateStudents() *NewStudents {
	db.NewRecord(s)
	db.Table("all_students").Create(&s)
	return s
}

func GetAllStudents() []Students {
	var Studs []Students
	//db.Find(&Studs)
	//db.Table("all_students").Select("class,firstname").Find(&Studs)
	db.Raw("select studentId,FirstName,LastName,Address,City,phoneNumber,schoolName,class from all_students").Scan(&Studs)
	//db.Table("studsubjmap").Select("studentIdMap,subjectIdMap").Scan(&Studs)
	return Studs

}

func GetStudentById(studentId int64) (*Students, *gorm.DB) {
	var getStudentById Students
	db := db.Table("all_students").Where("studentId=?", studentId).Find(&getStudentById)
	return &getStudentById, db
}

func DeleteStudentById(studentId int64) Students {
	var deletedStudent Students
	db.Table("all_students").Where("studentId=?", studentId).Delete(deletedStudent)
	return deletedStudent
}
