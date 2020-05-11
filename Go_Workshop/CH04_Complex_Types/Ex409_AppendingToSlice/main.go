package main

import (
	"fmt"
	"os"
)

func getPassArgs() [] string {
	var args [] string
	for i := 1; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	
	return args
}

func getLocals(extraLocals []string) [] string {
	var locales [] string
	locales = append(locales, "en-US", "fr-FR")
	locales = append(locales, extraLocals...)
	return locales
}

func main() {
	locales := getLocals(getPassArgs())
	fmt.Println("Locales to use:", locales)
}