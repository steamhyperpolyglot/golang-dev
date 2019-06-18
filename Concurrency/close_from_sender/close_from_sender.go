package main

import (
	"fmt"
	"time"
)

func main() {
	msg := make(chan string)
	// Adds an additional channel that indicates when we're finished.
	done := make(chan bool)
	until := time.After(5 * time.Second)

	// ch := make(chan bool)
	// timeout := time.After(600 * time.Millisecond)

	// Loops over a select with two channels and a default
	// go send(ch)
	go send(msg, done) // Pass two channels into send
	for {
		select {
		case m := <-msg: // If we get a message over our main channel, print something.
			fmt.Println(m)
		case <-until:
			done <- true // When we time-out, lets send know the process is done.
			time.Sleep(500 * time.Millisecond)
			return
			//default: // By default, we're going to sleep for a bit.
			//	println("*yawn*")
			//	time.Sleep(100 * time.Millisecond)
		}
	}
}

// ch is a receiving channel, while done is a sending channel.
func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
