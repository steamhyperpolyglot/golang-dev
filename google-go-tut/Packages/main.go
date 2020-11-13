package main

import "fmt"

// the package we just created
import "calculation"

func main() {
	x, y := 15, 10
	sum := calculation.Do_add(x, y)
	fmt.Println("Sum", sum)
}