package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The script has started at: ", start)
	fmt.Println("Saving the world...")
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println("The script has completed at: ", end)
	
	Day := time.Now().Weekday()
	Hour := time.Now().Hour()
	fmt.Println("Day: ", Day, "Hour: ", Hour)
	if Day.String() == "Wednesday" {
		if Hour >= 1 {
			fmt.Println("Performing full blown test!")
		} else {
			fmt.Println("Performing hit-n-run test!")
		}
	} else {
		fmt.Println("Performing hit-n-run test!")
	}
}