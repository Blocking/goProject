package main

import (
	"fmt"
	"math"
)

func main() {
	/*
	 导包形式 可以单个语句书写 也可以使用`()`进行整体书写
	*/
	fmt.Printf("now you have %g problems.", math.Nextafter(2, 3))
}
