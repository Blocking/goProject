package main

import (
	"fmt"
	"math/cmplx"
)

//变量打包声明
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	int1   int64      = 1<<63 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
	ii     int        = 9223372036854775807
)

func main() {
	const f = "%T(%v)\n"
	fmt.Printf(f, ToBe, ToBe)
	fmt.Printf(f, MaxInt, MaxInt)
	fmt.Printf(f, z, z)
	fmt.Printf(f, int1, int1)
	fmt.Println(ii)

}
