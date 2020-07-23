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

func GetTasks(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	log.Print("Get Tasks requested", id)
	fileContents := readFile("test.txt")
	if value, err := strconv.Atoi(id); err == nil {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents[(value - 1):(value * 10)]})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents})
	}
}

func readFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(bufio.NewReader(file))

	var tmp []string
	for scanner.Scan() {
		tmp = append(tmp, scanner.Text())
	}

	return tmp
}
