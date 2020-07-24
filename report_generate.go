package main

import (
	"encoding/csv"
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
	log.Print("create Report")
	total, completed, remining := calculatetotaltasks()
	file, err := os.OpenFile("report.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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

func Createreportperday(w http.ResponseWriter, r *http.Request) {
	Counttaskscompleted()

}
func maxtaskscompletedday(w http.ResponseWriter, r *http.Request) {
	Countmaxtaskscompleted()
}

func Counttaskscompleted() {
	input, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	counts := map[string]int{}

	for _, line := range lines {
		subline := strings.Split(string(line), " ")
		if len(subline) > 1 && len(subline[4]) > 1 {
			counts[subline[4]]++
		}
	}

	fmt.Println("counts: ", counts)
	log.Print("create Report")
	file, err := os.OpenFile("report.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	defer file.Close()

	if err != nil {
		os.Exit(1)
	}
	csvWriter := csv.NewWriter(file)

	for key, value := range counts {
		fmt.Println("Key:", key, "Value:", value)
		occurence := strconv.Itoa(value)
		x := []string{"Date", "Tasks"}
		y := []string{key, occurence}
		strWrite := [][]string{x, y}
		csvWriter.WriteAll(strWrite)
	}
	csvWriter.Flush()

}

func Countmaxtaskscompleted() {
	input, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")
	counts := map[string]int{}

	for _, line := range lines {
		subline := strings.Split(string(line), " ")
		if len(subline) > 1 && len(subline[4]) > 1 {
			counts[subline[4]]++
		}
	}

	fmt.Println("counts: ", counts)
	log.Print("create Report")
	file, err := os.OpenFile("report.csv", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
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
