package main

import (
	"log"
	"os"
	"fmt"
)

func simple(task int, logger *log.Logger) {
	for i := 0; i < 10; i++ {
		logger.Println(task, ":", i)
	}
}

func main() {
	logger := log.New(os.Stdout, "", 0)
	for i := 0; i < 10; i++ {
		go simple(i, logger)
	}

	// This is to prevent the program from closing before the printing has finished.
	var input string
	fmt.Scanln(&input)
}
