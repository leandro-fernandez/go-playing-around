package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Press Enter at anytime to finish...")

	messageChan := make(chan string)
	go sender(messageChan, 1)
	go sender(messageChan, 2)
	go sender(messageChan, 3)
	go sender(messageChan, 4)
	go receiver(messageChan)

	fmt.Scanln()
}

func receiver(message chan string) {
	for {
		msg := <-message
		fmt.Println(msg)
	}
}

func sender(message chan string, routineIndex int) {
	for {
		message <- fmt.Sprintf("a message from goroutine %d", routineIndex)
		time.Sleep(time.Duration(routineIndex) * time.Second)
	}
}
