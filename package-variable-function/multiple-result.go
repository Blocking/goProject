package main

import "fmt"

func swap(x, y string) (string, string) {
	return x, y
}
func main() {
	/**
	方法函数可以返回任意数量的返回值
	*/
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	fmt.Println(a)
	fmt.Println(b)
}
