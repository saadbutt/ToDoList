package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Delete Task with ID: " + id + " requested")
	if value, err := strconv.Atoi(id); err == nil {
		deleteFile("Files/test.txt", value)
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: "Task deleted"})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 400, Message: "Bad request", Response: "Id is not a number"})
	}
}

func deleteFile(filename string, tasknumber int) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, "{"+strconv.Itoa(tasknumber)) {
			lines[i] = ""
		}
	}
	fmt.Println("lines", lines)
	output := strings.Join(lines, "\n")
	fmt.Println("output", output)
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
