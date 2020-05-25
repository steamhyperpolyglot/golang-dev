package main

import (
	"fmt"
)

func fizzBuzz () {
	for i := 1; i <= 30; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fmt.Println("Main is in control")
	fizzBuzz()
	fmt.Println("Back to main")
}
