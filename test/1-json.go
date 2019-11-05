package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name   string `json:"name"`
	Age    int
	Emails []string
}

func main() {
	bingo := Person{
		Name:   "Bingo Huang",
		Age:    30,
		Emails: []string{"go@bingohuang.com", "me@bingohuang.com"},
	}
	bytes, e := json.Marshal(bingo)
	if e != nil {
		panic(e)
	}
	fmt.Printf("%s", bytes)

}
