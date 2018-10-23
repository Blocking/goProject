package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func add1(x, y int) int {
	return x + y
}
func main() {
	/*
		函数可以接受 多个参数 或者没有参数
		变量类型 跟在变量名之后
	*/
	fmt.Println(add(42, 13))
	/**
	当俩个或者多个函数命名参数是统一类型 则除了最后一个参数的类型外 其它的参数类型可以省略
	*/
	fmt.Println(add1(42, 13))
}
