package main

import (
	"fmt"
	"time"
)

func marcoPoloPrinter(receiver <-chan string, sender chan<- string, msg string) {
	// Loop endlessly
	for {
		// Wait here until we receive a message
		m := <-receiver
		// Print that message
		fmt.Println(m)
		// Send our message
		sender <- msg
		// Slow this down a bit
		time.Sleep(time.Second * 1)
	}
}

func main() {
	// Make a channel for our polo message
	polo := make(chan string)
	// Make a channel for our marco message
	marco := make(chan string)

	// Set up to concurrent go routines to communicate
	go marcoPoloPrinter(marco, polo, "marco")
	go marcoPoloPrinter(polo, marco, "polo")

	// Trigger the start by sending the first message
	marco <- "marco"

	// So we can end at any time
	var input string
	fmt.Scanln(&input)
}
