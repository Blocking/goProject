package main

import "fmt"

func add(x int, y int) int {
	return x + y
}
func main() {
	/*
		函数可以接受 多个参数 或者没有参数
		变量类型 跟在变量名之后
	*/
	fmt.Println(add(42, 13))
}
