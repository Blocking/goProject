package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("hello world 世界")
	fmt.Println("我的幸运数字：", rand.Int())
	fmt.Printf("现在你有 %g 问题了", math.Sqrt(4))
	fmt.Println(math.Pi)
	fmt.Println(add(12, 45))
	fmt.Println(add2(12, 45))
	fmt.Println(swap("12", "45"))
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	c, d := split(17)
	fmt.Println(c, d)
}
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func swap(x, y string) (string, string) {
	return y, x
}

func add2(x, y int) int {
	return x + y
}

func add(x int, y int) int {
	return x + y
}
