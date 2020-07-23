package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func GetTasks(w http.ResponseWriter, r *http.Request) {
	log.Print("Get Tasks requested")
	fileContents := readFile("test.txt")
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents})
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
