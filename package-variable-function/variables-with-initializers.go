package main

import "fmt"

var i, j int = 1, 2

func main() {
	/**
	变量定义可以包含初始值 多个变量初始化时需一一对应
	倘若使用表达式进行初始化赋值 则可以省略类型；变量从表达式的初始值中获取类型
	*/
	var c, python, java = true, false, "no!"
	var m, n = 1 + 4, "dfsd" + "4565"
	fmt.Println(m, n)
	fmt.Println(i, j, c, python, java)
}
