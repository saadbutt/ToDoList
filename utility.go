package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateMockData() {
	// tasks = append(tasks, Task{ID: 1, Name: "Create a REST API", Status: Ongoing, StartDate: "20/07/2018"})
	// tasks = append(tasks, Task{ID: 2, Name: "Create a Client in IONIC", Status: Pending, StartDate: "27/07/2018"})
	//	tasks = append(tasks, Task{ID: 3, Name: "Publish the result", Status: Blocked})
}

const ID = 0

/*
 * Enum that defines the status of task
 */
type TaskStatus int

const (
	Pending  TaskStatus = 0
	Ongoing  TaskStatus = 1
	Done     TaskStatus = 2
	Blocked  TaskStatus = 3
	Rejected TaskStatus = 4
)

/*
 * Struct that defines a task
 */
type Task struct {
	ID             int    `json:"id"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Created        string `json:"created"`
	DueDate        string `json:"duedate"`
	CompletionDate string `json:"completiondate"`
	Completed      bool   `json:"completed"`
}

//file attachments (e.g. an image)

var tasks []Task

type CustomResponse struct {
	HttpCode int
	Message  string
	Response interface{}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Get Task with ID: " + id + " requested")
	if value, err := strconv.Atoi(id); err == nil {
		for _, item := range tasks {
			if item.ID == value {
				json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: item})
				return
			}
		}
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 404, Message: "Not Found", Response: "No task found with id: " + id})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 400, Message: "Bad request", Response: "Id is not a number"})
	}
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = getLastID()
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 201, Message: "Created", Response: task})
	log.Print("Created Task with ID: " + strconv.Itoa(task.ID))
}

func writeFileandCallAPI(filename string, lines Task) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666) //os.Create(filename) //example of multiple results from a function where one is the error code
	if err != nil {
		panic("could not open todo file")
	}
	defer file.Close() //will call file's close function at the end of writeFile

	w := bufio.NewWriter(file)
	defer w.Flush() //interesting, two deferred funcs, one needs to be called first....

	// for _, each := range lines { //ignore the first param with "_"
	// 	fmt.Fprint(w, each)
	// }
}

func getLastID() int {
	var task Task
	if len(tasks) == 0 {
		return 1
	}
	task = tasks[len(tasks)-1]
	return task.ID + 1
}
