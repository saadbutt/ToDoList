package main

import (
	"bufio"
	"encoding/json"
	"fmt"
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
	fileContents := readFile("Files/test.txt")
	if value, err := strconv.Atoi(id); err == nil {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents[(value - 1):(value * 10)]})
	} else {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents})
	}
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	limit, ok := r.URL.Query()["limit"]
	offset, okay := r.URL.Query()["offset"]

	fileContents := readFile("Files/test.txt")
	fmt.Println("fileContents:", len(fileContents))
	if !ok || len(limit[0]) < 0 || !okay || len(offset[0]) < 0 {
		log.Println("Url Param 'key' is missing")
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents})

	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key, _ := strconv.Atoi(limit[0])
	value, _ := strconv.Atoi(offset[0])
	log.Println("Url Param 'key' is: "+string(key), string(value))

	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: paginate(fileContents, value, key)})

}

func paginate(x []string, skip int, size int) []string {
	limit := func() int {
		if skip+size > len(x) {
			return len(x)
		} else {
			return skip + size
		}

	}

	start := func() int {
		if skip > len(x) {
			return len(x)
		} else {
			return skip
		}

	}
	return x[start():limit()]
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
