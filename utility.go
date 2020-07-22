package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"

	"github.com/gorilla/mux"
)

func CreateMockData() {
	tasks = append(tasks, Task{ID: 1, Name: "Create a REST API", Status: Ongoing, StartDate: "20/07/2018"})
	tasks = append(tasks, Task{ID: 2, Name: "Create a Client in IONIC", Status: Pending, StartDate: "27/07/2018"})
	tasks = append(tasks, Task{ID: 3, Name: "Publish the result", Status: Blocked})
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
	ID        int
	Name      string
	Status    TaskStatus
	StartDate string
	EndDate   string
}

var tasks []Task

type CustomResponse struct {
	HttpCode int
	Message  string
	Response interface{}
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Print("Get Tasks requested")
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: tasks})
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

func writeFileandCallAPI(filename string, lines []string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666) //os.Create(filename) //example of multiple results from a function where one is the error code
	if err != nil {
		panic("could not open todo file")
	}
	defer file.Close() //will call file's close function at the end of writeFile

	w := bufio.NewWriter(file)
	defer w.Flush() //interesting, two deferred funcs, one needs to be called first....

	for _, each := range lines { //ignore the first param with "_"
		fmt.Fprint(w, each+"\n")
	}
}

func CreateTasktest(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = getLastID()
	tasks = append(tasks, task)
	log.Print("tasks", tasks)
	log.Print("reflect", reflect.ValueOf(tasks))
	args := []string{tasks[0].Name, tasks[1].Name, tasks[2].Name}

	writeFile("test.txt", args)
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 201, Message: "Created", Response: task})
	log.Print("Created Task with ID: " + strconv.Itoa(task.ID))
}

func getLastID() int {
	var task Task
	task = tasks[len(tasks)-1]
	return task.ID + 1
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Delete Task with ID: " + id + " requested")
	if value, err := strconv.Atoi(id); err == nil {
		for index, item := range tasks {
			if item.ID == value {
				tasks = tasks[:index]
				json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: "Task deleted"})
				return
			}
		}
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 404, Message: "Not Found", Response: "No task found with id: " + id})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 400, Message: "Bad request", Response: "Id is not a number"})
	}
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Update Task with ID: " + id + " requested")
	if value, err := strconv.Atoi(id); err == nil {
		var task Task
		_ = json.NewDecoder(r.Body).Decode(&task)
		for index, item := range tasks {
			if item.ID == value {
				tasks[index].Name = task.Name
				tasks[index].Status = task.Status
				tasks[index].StartDate = task.StartDate
				tasks[index].EndDate = task.EndDate
				json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: tasks[index]})
				return
			}
		}
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 404, Message: "Not Found", Response: "No task found with id: " + id})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 400, Message: "Bad request", Response: "Id is not a number"})
	}
}
