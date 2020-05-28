package main

import "fmt"

func main() {
	fname := "Joe"
	gpa := 3.75
	hasJob := true
	hourlyWage := 45.53
	age := 24
	fmt.Printf("%s has a gps of %f.\n", fname, gpa)
	fmt.Printf("He has a job equals %t.\n", hasJob)
	fmt.Printf("He is %d earning %v per hour.\n", age, hourlyWage)
	
	fmt.Println("Formatting of floating-point literals:")
	v := 1234.0
	v1 := 1234.6
	v2 := 1234.67
	v3 := 1234.678
	v4 := 1234.6789
	v5 := 1234.67891
	
	fmt.Printf("%-10.0f\n", v)
	fmt.Printf("%-10.1f\n", v1)
	fmt.Printf("%-10.2f\n", v2)
	fmt.Printf("%-10.3f\n", v3)
	fmt.Printf("%-10.4f\n", v4)
	fmt.Printf("%-10.5f\n", v5)
}