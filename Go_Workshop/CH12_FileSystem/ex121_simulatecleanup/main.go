package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool)

	signal.Notify(sigs, syscall.SIGINT)

	go func() {
		s := <-sigs
		switch s {
		case syscall.SIGINT:
			fmt.Println()
			fmt.Println("My process has been interrupted. Someone might've pressed CTRL-C.")
			fmt.Println("Some clean up is occuring")
			cleanUp()
			done <- true
		}
	}()
	fmt.Println("Program is blocked until a signal is caught (ctrl-c)")
	<-done
	fmt.Println("Out of here")
}

func cleanUp() {
	fmt.Println("Simulating clean up")
	for i := 0; i <= 10; i++ {
		fmt.Println("Deleting Files... Not really.", i)
		time.Sleep(1 * time.Second)
	}
}
