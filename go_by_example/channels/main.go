package main

import "fmt"

func main() {
	// create a new channel
	messages := make(chan string)

	// send "ping" to the message
	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)

}
