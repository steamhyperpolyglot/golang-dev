package main

import (
	"log"
)

func main() {
	name := "Thanos"
	log.Println("Demo app")
	log.Printf("%s is here!", name)
	log.Print("Run")
	
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)

	log.Println("Demo app")
	log.Printf("%s is here!", name)
	log.Print("Run")
}