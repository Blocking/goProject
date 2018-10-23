package main

import "fmt"

func split(sum int) (x, y, m int) {
	x = sum * 4 / 9
	y = sum - x
	return y, x, 1
}

func main() {
	/**
	函数方法的返回值可以当做变量使用 可以被赋值
	没有参数的 return语句 直接返回当前 x,y的值 通常不建议这么做 影响代码可读性
	*/
	fmt.Println(split(17))
}
