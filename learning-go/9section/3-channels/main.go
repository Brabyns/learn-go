package main

import (
	"fmt"
	"time"
)

func main() {
	messages := make(chan string) //Unbuffered channel

	go func() {
		fmt.Println("Sending a message to messages channel")
		messages <- "Helllo from messages channel"
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("About to get a message from channel")
	msg := <-messages
	fmt.Println(msg)
}