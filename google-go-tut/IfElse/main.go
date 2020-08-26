package main

import "fmt"

func main() {
	var x = 100
	if x < 10 {
		// Executes if x is less than 10
		fmt.Println("x is less than 10")
	} else if x >= 10 && x <= 90 {
		// Executes if x >= 10 and x <= 90
		fmt.Println("x is between 10 and 90")
	} else {
		// Executes if both above cases fail i.e. x > 90
		fmt.Println("x is greater than 90")
	}
}