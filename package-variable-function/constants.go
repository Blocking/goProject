package main

import "fmt"

const Pi = 3.14

func main() {
	//常量定义使用`const`关键字 来声明 不能使用`:=` ；常量可以是字符、字符串、布尔或者数字类型的值
	const world = "世界"
	fmt.Println("hello", world)
	fmt.Println("Happy", Pi, "day")

	const Truth = true
	fmt.Println("GO rules?", Truth)
}
