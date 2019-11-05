package main

import "fmt"

func main() {
	sum := 1
	//此处 `for` 与 `java`中的`while` 用法作用相同
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}
