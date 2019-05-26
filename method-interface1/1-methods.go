package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 当 v 是 Vertex 的时候 Scale 方法没有任何作用。`Scale` 修改 `v`。当 v 是一个值（非指针），方法看到的是 Vertex 的副本，并且无法修改原始值。
//Abs 的工作方式是一样的。只不过，仅仅读取 `v`。所以读取的是原始值（通过指针）还是那个值的副本并没有关系。
func (v Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type Myfloat float64

func (f Myfloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Println(v.Abs())

	f := Myfloat(-math.Sqrt2)
	fmt.Println(f.Abs())

	v.Scale(5)
	fmt.Println(v, v.Abs())
}
