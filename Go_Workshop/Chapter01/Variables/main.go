package main

import (
	"fmt"
	"time"
)

var (
	Debug = false
	LogLevel = "info"
	startUpTime = time.Now()
)

// var foo string = "bar"

func main() {
	fmt.Println(Debug, LogLevel, startUpTime)
}
