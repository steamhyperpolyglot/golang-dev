package main

import (
	"encoding/json"
	"fmt"
)

type book struct {
	ISBN            string  `json:"isbn"`
	Title           string  `json:"title"`
	YearPublished   int     `json:"yearpub,omitempty"`
	Author          string  `json:"author"`
	CoAuthor        string  `json:"coauthor,omitempty"`
}

func main() {
	var b book
	b.ISBN = "9933HIST"
	b.Title = "Greatest of all Books"
	b.Author = "John Adams"
	json, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s", json)
}