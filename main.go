package main

import (
	"fmt"
	"time"
)

func main() {
	messageChan := make(chan string)
	go sender(messageChan)
	go receiver(messageChan)

	fmt.Println("Press Enter to finish....")
	fmt.Scanln()
}

func receiver(message chan string) {
	for {
		msg := <-message
		fmt.Println(msg)
	}
}

func sender(message chan string) {
	for {
		message <- "a message"
		time.Sleep(2 * time.Second)
	}
}
