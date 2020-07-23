package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

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
	fmt.Println("lines", lines)
	output := strings.Join(lines, "\n")
	fmt.Println("output", output)
	err = ioutil.WriteFile(filename, []byte(output), 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
