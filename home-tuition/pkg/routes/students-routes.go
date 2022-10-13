package routes

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rebontadeb/home-tuition/pkg/controllers"
)

var StudentRoutes = func(router *mux.Router) {

	router.HandleFunc("/allStudents", controllers.GetAllStudents).Methods("GET")
	router.HandleFunc("/allStudents/{studentId}", controllers.GetAllStudentsById).Methods("GET")
	router.HandleFunc("/allStudents", controllers.OnboardStudent).Methods("POST")
	router.HandleFunc("/allStudents/{studentId}", controllers.DeleteStudentById).Methods("DELETE")
}
