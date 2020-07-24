package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	Logger("Update Task with ID: " + id + " requested")
	log.Print("Update Task with ID: " + id + " requested")
	if value, err := strconv.Atoi(id); err == nil {
		var task Task
		_ = json.NewDecoder(r.Body).Decode(&task)
		updateFile("Files/test.txt", value, task)

		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: value})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 400, Message: "Bad request", Response: "Id is not a number"})
	}
}

func updateFile(filename string, tasknumber int, task Task) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "{"+strconv.Itoa(tasknumber)) {
			lines[i] = "{" + strconv.Itoa(tasknumber) + " " + task.Title + " " + task.Description + " " + task.Created + " " + task.CompletionDate + " " + task.DueDate + " " + strconv.FormatBool(task.Completed) + "}"
		}
	}
	output := strings.Join(lines, "\n")
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
