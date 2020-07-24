package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

// func makeAdd(filename string) *Tasks {

// 	w, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	defer w.Close()
// 	task := strings.Join(args, " ")
// 	_, err = fmt.Fprintln(w, task)
// 	fmt.Printf("Task added: %s\n", task)
// 	return err
// }
func writeFile(filename string, lines Task) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666) //os.Create(filename) //example of multiple results from a function where one is the error code
	if err != nil {
		panic("could not open todo file")
	}
	defer file.Close() //will call file's close function at the end of writeFile

	w := bufio.NewWriter(file)
	defer w.Flush() //interesting, two deferred funcs, one needs to be called first....

	fmt.Fprintln(w, lines)
}

func CreateTasktest(w http.ResponseWriter, r *http.Request) {
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = getLastID()
	tasks = append(tasks, task)
	log.Print("tasks", tasks)
	//args := []string{tasks[0].Name, tasks[1].Name, tasks[2].Name}
	writeFile("Files/test.txt", task)
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 201, Message: "Created", Response: task})
	log.Print("Created Task with ID: " + strconv.Itoa(task.ID))
}
