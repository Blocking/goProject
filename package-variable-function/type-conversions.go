package main

import (
	"fmt"
	"math"
)

func main() {
	//go语言类型转换需要显示转换
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z int = int(f)
	fmt.Println(x, y, z)

}
