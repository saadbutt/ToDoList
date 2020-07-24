package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Count of total tasks, completed tasks, and remaining tasks (aggregate all 3 in parallel)
func Createreport(w http.ResponseWriter, r *http.Request) {
	Logger("create Report of Count of total tasks, completed tasks, and remaining tasks")
	total, completed, remining := calculatetotaltasks()
	file, err := os.Create("Files/tasksreport.csv")
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}

	x := []string{"Total", "Completed", "Remaining"}
	y := []string{total, completed, remining}
	csvWriter := csv.NewWriter(file)
	strWrite := [][]string{x, y}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: "PDF Generated. Downloadable Link http://localhost:8000/Files/tasksreport.csv"})

}

func calculatetotaltasks() (string, string, string) {
	input, err := ioutil.ReadFile("Files/test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	counttotaltasks := 0
	completedtasks := 0
	for _, line := range lines {
		if line != "" {
			counttotaltasks += 1
		}
		if strings.Contains(line, "true}") {
			completedtasks += 1
		}
	}
	log.Println("completedtasks", completedtasks)
	log.Println("total tasks", counttotaltasks)
	log.Print("Remaining tasks", counttotaltasks-completedtasks)
	return strconv.Itoa(counttotaltasks), strconv.Itoa(completedtasks), strconv.Itoa((counttotaltasks - completedtasks))
}

//Average number of tasks completed per day (aggregate average in parallel for each day)
func createReportPerDay(w http.ResponseWriter, r *http.Request) {
	Logger("create Report of Average number of tasks completed per day")
	Counttaskscompleted()
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: "PDF Generated. Downloadable Link http://localhost:8000/Files/CompletedTaskreport.csv"})

}

// maximum number of tasks were completed in a single day
func maxTasksCompleted(w http.ResponseWriter, r *http.Request) {
	Logger("create Report of maximum number of tasks were completed in a single day")
	Countmaxtaskscompleted()
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: "PDF Generated. Downloadable Link http://localhost:8000/Files/maxtasksreport.csv"})

}

func maxTasksAdded(w http.ResponseWriter, r *http.Request) {
	Logger("create Report of maxTasksAdded")
	Counttasksadded()
	json.NewEncoder(w).Encode(&CustomResponse{HttpCode: 200, Message: "OK", Response: "PDF Generated.  Downloadable Link http://localhost:8000/Files/addedtasksreport.csv"})

}

func Counttaskscompleted() {
	input, err := ioutil.ReadFile("Files/test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	counts := map[string]int{}
	totaltasks := 0
	for _, line := range lines {
		subline := strings.Split(string(line), " ")
		if len(subline) > 1 && len(subline[4]) > 1 {
			totaltasks += 1
			counts[subline[4]]++
		}
	}

	log.Print("create Report")
	file, err := os.Create("Files/CompletedTaskreport.csv")
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)

	for key, value := range counts {
		fmt.Println("Key:", key, "Value:", value)
		occurence := strconv.Itoa(value / totaltasks)
		x := []string{"Date", "AVGTasks"}
		y := []string{key, occurence}
		strWrite := [][]string{x, y}
		csvWriter.WriteAll(strWrite)
	}
	csvWriter.Flush()

}

func Countmaxtaskscompleted() {
	input, err := ioutil.ReadFile("Files/test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	counts := map[string]int{}

	for _, line := range lines {
		subline := strings.Split(string(line), " ")
		if len(subline) > 1 && len(subline[4]) > 1 && strings.Contains(line, "true}") {
			counts[subline[4]]++
		}
	}

	log.Print("create Report")
	file, err := os.Create("Files/maxtasksreport.csv")
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)
	max := 0
	var date string
	for key, value := range counts {
		fmt.Println("Key:", key, "Value:", value)
		if max < value {
			date = key
			max = value
		}
	}

	occurence := strconv.Itoa(max)
	x := []string{"Date", "MAX Tasks"}
	y := []string{date, occurence}
	strWrite := [][]string{x, y}
	csvWriter.WriteAll(strWrite)
	csvWriter.Flush()

}

func Counttasksadded() {
	input, err := ioutil.ReadFile("Files/test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	counts := map[string]int{}
	maxdays := 0
	for _, line := range lines {
		subline := strings.Split(string(line), " ")
		if len(subline) > 1 && len(subline[3]) > 1 {
			counts[subline[3]]++

			fmt.Println(subline[3], counts[subline[3]])
			if maxdays < counts[subline[3]] {
				maxdays = counts[subline[3]]
			}
		}
	}

	log.Print("create Report")
	file, err := os.Create("Files/addedtasksreport.csv")
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)
	x := []string{"Date", "MAX Tasks"}

	for key, value := range counts {
		fmt.Println("Key:", key, "Value:", value)
		if maxdays == value {
			y := []string{key, strconv.Itoa(value)}
			strWrite := [][]string{x, y}
			csvWriter.WriteAll(strWrite)
		}
	}

	csvWriter.Flush()
}
