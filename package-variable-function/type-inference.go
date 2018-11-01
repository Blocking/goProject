package main

import "fmt"

func main() {
	//类型推导 由初始值反推变量类型
	v := 42
	fmt.Printf("v is of type %T\n", v)
}
