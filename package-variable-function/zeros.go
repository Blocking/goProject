package main

import "fmt"

func main() {
	//倘若不赋值 数值类型的默认值是`0`  布尔类型的默认值为`false` 字符串的默认值为`""`
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q \n", i, f, b, s)
}
