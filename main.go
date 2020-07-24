package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Tasks todo tasks struct
type Tasks struct {
	Title          string `json:"title"`
	Description    string `json:"description"`
	Created        string `json:"created"`
	DueDate        string `json:"duedate"`
	CompletionDate string `json:"completiondate"`
	Completed      bool   `json:"completed"`
	Public         bool   `json:"public"`
	Allday         bool   `json:"allday"`
}

type FileMD struct {
	fileName string
	fileSize int
}

func main() {

	var dir = "./Files"
	//	fs := http.FileServer(http.Dir("./swaggerui/"))

	CreateMockData()
	router := mux.NewRouter()
	//	router.HandleFunc("/task", GetTasks).Methods("GET")
	router.HandleFunc("/generatereport", Createreport).Methods("GET")
	router.HandleFunc("/task", GetTask).Methods("GET")
	router.HandleFunc("/task", CreateTasktest).Methods("POST")
	router.HandleFunc("/task/{id}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", UpdateTask).Methods("PUT")
	router.HandleFunc("/createReportPerDay", createReportPerDay).Methods("GET")
	router.HandleFunc("/maxTasksCompleted", maxTasksCompleted).Methods("GET")
	router.HandleFunc("/maxTasksAdded", maxTasksAdded).Methods("GET")
	router.PathPrefix("/Files/").Handler(http.StripPrefix("/Files/", http.FileServer(http.Dir(dir))))
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
	// router.PathPrefix("/swaggerui/").Handler(http.StripPrefix("/swaggerui/", fs))
	log.Fatal(http.ListenAndServe(":8000", router))

}
