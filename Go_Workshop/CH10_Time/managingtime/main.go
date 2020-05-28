package main

import (
	"fmt"
	"time"
)

func main() {
	TimeToManipulate := time.Now()
	ToBeAdded := time.Duration(10 * time.Second)
	fmt.Println("The original time: ", TimeToManipulate)
	fmt.Println(ToBeAdded, " duration later:", TimeToManipulate.Add(ToBeAdded))
}