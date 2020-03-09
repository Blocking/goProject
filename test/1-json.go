package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Person struct {
	Name   string `json:"name"`
	Age    int
	Emails []string
	Timer  time.Time
}

func main() {
	bingo := &Person{
		Name:   "Bingo Huang",
		Age:    30,
		Emails: []string{"go@bingohuang.com", "me@bingohuang.com"},
		Timer:  time.Now(),
	}

	bytes, e := json.Marshal(bingo)
	if e != nil {
		panic(e)
	}
	fmt.Println(string(bytes))

}
