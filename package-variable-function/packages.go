package main

import (
	"fmt"
	"math/rand"
)

func main() {
	/*
		程序入口包是main 当前程序使用了 fmt包 以及math/rand 包
		按照惯例 包名与导入路径的最后一个目录一致。例如 math/rand  包名为 rand 所以使用 rand.Intn()调用
	*/
	fmt.Println("My favorite number is", rand.Intn(10))
}
