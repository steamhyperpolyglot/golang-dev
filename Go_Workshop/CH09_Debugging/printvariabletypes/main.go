package main

import (
	"fmt"
)

type person struct {
	lname string
	age int
	salary float64
}

func main() {
	fname := "Joe"
	grades := []int {100, 87, 67}
	states := map[string]string {
		"KY": "Kentucky", "WV": "West Virginia", "VA": "Virginia" }
	p := person { lname: "Lincoln", age: 210, salary: 25000.00 }
	fmt.Printf("fname is of type %T\n", fname)
	fmt.Printf("grades is of type %T\n", grades)
	fmt.Printf("states is of type %T\n", states)
	fmt.Printf("p is of type %T\n", p)
}