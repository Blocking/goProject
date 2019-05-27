package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn`t not work",
	}
}

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}

	i, err1 := strconv.Atoi("42")
	if err1 != nil {
		fmt.Printf("couldn`t convert number: %v\n", err1)
	}
	fmt.Println("Converted integer:", i)
}
