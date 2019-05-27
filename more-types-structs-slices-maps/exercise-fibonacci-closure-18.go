package main

import "fmt"

/*
练习：斐波纳契闭包
现在来通过函数做些有趣的事情。

实现一个 fibonacci 函数，返回一个函数（一个闭包）可以返回连续的斐波纳契数。
*/
func fibonacci() func() int {
	var count, sum1, sum2 = 0, 1, 1
	return func() int {
		switch count {
		case 0, 1:
			count++
		default:
			count++
			sum1, sum2 = sum2, sum1+sum2
		}
		return sum2
	}
}
func fibonacci2() func() int {
	var x, y = -1, 1
	return func() int {
		x, y = y, x+y
		return y
	}
}

func main() {
	f := fibonacci2()
	for i := 1; i < 10; i++ {
		fmt.Println(f())
	}
}
