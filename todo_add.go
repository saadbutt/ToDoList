package main

import (
	"bufio"
	"fmt"
	"os"
)

// func makeAdd(filename string) *Tasks {

// 	w, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
// 	if err != nil {
// 		return err
// 	}
// 	defer w.Close()
// 	task := strings.Join(args, " ")
// 	_, err = fmt.Fprintln(w, task)
// 	fmt.Printf("Task added: %s\n", task)
// 	return err
// }
func writeFile(filename string, lines []string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666) //os.Create(filename) //example of multiple results from a function where one is the error code
	if err != nil {
		panic("could not open todo file")
	}
	defer file.Close() //will call file's close function at the end of writeFile

	w := bufio.NewWriter(file)
	defer w.Flush() //interesting, two deferred funcs, one needs to be called first....

	for _, each := range lines { //ignore the first param with "_"
		fmt.Fprint(w, each+"\n")
	}
}
