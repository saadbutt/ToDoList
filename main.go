package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

	args := []string{"a", "b", "c"}

	writeFile("test.txt", args)

	CreateMockData()
	router := mux.NewRouter()
	router.HandleFunc("/task", GetTasks).Methods("GET")
	router.HandleFunc("/task/{id}", GetTask).Methods("GET")
	router.HandleFunc("/task", CreateTasktest).Methods("POST")
	router.HandleFunc("/task/{id}", DeleteTask).Methods("DELETE")
	router.HandleFunc("/task/{id}", UpdateTask).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", router))
}
