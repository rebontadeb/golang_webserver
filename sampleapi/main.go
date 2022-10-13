package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article

//ALL ARTICLES
func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Anaconda", Desc: "Movie Hollywood", Content: "Movie about snakes"},
		Article{Title: "Home Alone", Desc: "Movie Hollywood", Content: "Movie about a child"},
	}
	fmt.Println("Endpoints Hit : All Articles Endpoints")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(articles)
}

//UPDTAE_STUDENTS
func updateStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "UPDATE STUDENT")
	fmt.Println("Endpoints Hit : All Articles Endpoints")
}

//GET_STUDENTS_BY_ID
func getStudentsByID(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET ALL STUDENTS BY ID")
	fmt.Println("Endpoints Hit : All Articles Endpoints")
}

//GET_STUDENTS
func getStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET ALL STUDENTS")
	fmt.Println("Endpoints Hit : All Articles Endpoints")
}

//ADD_STUDENTS
func addStudents(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ADD STUDENTS")
}

func allSubjects(w http.ResponseWriter, r *http.Request) {

	type subject struct {
		SubjectId int    `json:"SubjectId"`
		subject   string `json:"subject"`
	}
	var subj subject
	type subjects []subject
	var allSubj []subject
	//++++++++++++++++++++++++++CONNECTION++++++++++++++++++++++++++
	db, err := sql.Open("mysql", "class:class@tcp(127.0.0.1:33061)/ht_academics")
	if err != nil {
		log.Fatal("Error while opening database connection:", err.Error())
	}
	defer db.Close()
	//++++++++++++++++++++++++++CONNECTION++++++++++++++++++++++++++
	//++++++++++++++++++++++++++QUERY++++++++++++++++++++++++++
	results, err := db.Query("SELECT SubjectId, subject FROM ht_academics.all_subjects")
	if err != nil {
		log.Fatal("Query failed:", err.Error())
	}
	defer results.Close()
	//++++++++++++++++++++++++++QUERY++++++++++++++++++++++++++

	for results.Next() {
		err = results.Scan(&subj.SubjectId, &subj.subject)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		log.Println(subj.SubjectId, subj.subject)
		jsonData := subjects{subject{SubjectId: subj.SubjectId, subject: subj.subject}}
		fmt.Println(jsonData)
		//allSubj = append(allSubj, string(jsonData))
		//log.Printf("The type is %T ", allSubj)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allSubj)
}

func handleRequest() {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/api/allArticles", allArticles).Methods("GET")
	myRouter.HandleFunc("/api/allSubjects", allSubjects).Methods("GET")
	myRouter.HandleFunc("/api/addStudents", addStudents).Methods("POST")
	myRouter.HandleFunc("/api/getStudents", getStudents).Methods("GET")
	myRouter.HandleFunc("/api/getStudents/{id}", getStudentsByID).Methods("GET")
	myRouter.HandleFunc("/api/updateStudents/{id}", updateStudents).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main() {
	handleRequest()
}
