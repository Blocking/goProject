package main

import (
	"strconv"
)

var c, python, java bool

func main() {
	number1, _ := strconv.Atoi("123")
	number2, _ := strconv.Atoi("323")
	println("Sum:", number1+number2)
}
