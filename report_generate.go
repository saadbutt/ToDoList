package main

import (
	"encoding/csv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Count of total tasks, completed tasks, and remaining tasks (aggregate all 3 in parallel)
func Createreport(w http.ResponseWriter, r *http.Request) {
	log.Print("create Report")
	total, completed, remining := calculatetotaltasks()
	file, err := os.OpenFile("report.csv", os.O_CREATE|os.O_WRONLY, 0777)
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

}
func calculatetotaltasks() (string, string, string) {
	input, err := ioutil.ReadFile("test.txt")
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

//- Average number of tasks completed per day (aggregate average in parallel for each day)
func Createreportperday(w http.ResponseWriter, r *http.Request) {

}
