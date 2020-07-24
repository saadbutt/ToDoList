package main

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
)

func GetTask(w http.ResponseWriter, r *http.Request) {
	log.Print("GET TASK REQUEST")
	Logger("GET TASK REQUEST")
	limit, ok := r.URL.Query()["limit"]
	offset, okay := r.URL.Query()["offset"]

	fileContents := readFile("Files/database.txt")
	if !ok || len(limit[0]) < 0 || !okay || len(offset[0]) < 0 {
		json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: fileContents})

	}

	// Query()["key"] will return an array of items,
	// we only want the single item.
	key, _ := strconv.Atoi(limit[0])
	value, _ := strconv.Atoi(offset[0])

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
