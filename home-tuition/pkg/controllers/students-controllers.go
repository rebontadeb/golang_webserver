package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rebontadeb/home-tuition/pkg/models"
	"github.com/rebontadeb/home-tuition/pkg/utils"
)

var NewStudent models.Students

func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	myStudent := models.GetAllStudents()
	res, _ := json.Marshal(myStudent)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetAllStudentsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	studId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	studentDetailsById, _ := models.GetStudentById(studId)
	res, _ := json.Marshal(studentDetailsById)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func OnboardStudent(w http.ResponseWriter, r *http.Request) {
	OnboardStudent := &models.NewStudents{}
	utils.ParseBody(r, OnboardStudent)
	s := OnboardStudent.CreateStudents()
	res, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStudentById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	studentId := vars["studentId"]
	studId, err := strconv.ParseInt(studentId, 0, 0)
	if err != nil {
		fmt.Println("Error While Parsing")
	}
	student := models.DeleteStudentById(studId)
	res, _ := json.Marshal(student)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
