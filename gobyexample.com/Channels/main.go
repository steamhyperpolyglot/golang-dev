package main

import "fmt"

func main() {
	
	messages := make (chan string)
	
	// Semd a value into a channel using the channel <- syntax.
	// Here we send "ping" to the messages channel we made above,
	// from a new goroutine.
	go func() { messages <- "ping" }()
	
	msg := <- messages
	fmt.Println(msg)
}