package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	time.Sleep(time.Second)
	duration := time.Since(start)
	fmt.Println("It took time.Sleep(1)", duration, "to finish.")
}