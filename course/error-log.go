package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

type employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found!")

func errorLog() {
	//error1()
	//log.SetPrefix("main():")

	//log.Print("hey,I'm an log!")
	//log.Fatal("Hey,I'm a log!")
	//logt()
	//log.Panic("Hey,I'm an error log! ")
	//fmt.Print("Can you see me?")
	//logFile()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log.Print("Hey! I'm a log message!")

	log.Debug().Int("EmployeeID", 1001).Msg("Getting employee information")
	log.Debug().
		Str("Name", "John").
		Send()
}

func logFile() {
	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		//log.Fatal(err)
	}

	defer file.Close()

	//log.SetOutput(file)
	log.Print("Hey, I'm a log!")
}

func logt() {
	log.Print("hey,I'm an func!")
}

func error1() {
	information, err := getInfomation(1001)

	if errors.Is(err, ErrNotFound) {
		fmt.Printf("NOT FOUND: %v\n", err)
	} else {
		fmt.Println(information)
	}
}

func getInfomation(id int) (*employee, error) {
	if id == 1001 {
		return nil, ErrNotFound
	}
	callEmployee, err := apiCallEmployee(1000)
	if err != nil {
		return nil, fmt.Errorf("Got an error when getting the employee information: %v", err)
	}
	return callEmployee, err
}
func apiCallEmployee(id int) (*employee, error) {
	emp := employee{LastName: "Doe", FirstName: "John"}
	return &emp, nil
}
