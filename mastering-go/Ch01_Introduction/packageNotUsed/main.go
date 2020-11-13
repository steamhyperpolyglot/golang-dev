package main

import (
	"fmt"
	_ "os"      // Adding an underscore in front of the package will help to avoid generating an error
)

func main() {
	fmt.Println("Hello there!")
}