package main

import (
	"fmt"
	"time"
)

func main() {
	Debug := false
	LogLevel := "info"
	startUpTime := time.Now()

	// Doing it in a single line
	// Debug, LogLevel, startUpTime := false, "info", time.Now()

	fmt.Println(Debug, LogLevel, startUpTime)
}
