package main

import (
	"fmt"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "Hello, my name is Sherman Chen")
}

func main() {
	http.HandleFunc("/", hello)
	http.ListenAndServe("localhost:5000", nil)
}
