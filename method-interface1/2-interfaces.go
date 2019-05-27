package main

import (
	"fmt"
	"math"
)

/*
接口
接口类型是由一组方法定义的集合。

接口类型的值可以存放实现这些方法的任何值。

注意： 列子代码的 29 行存在一个错误。 由于 Abs 只定义在 *Vertex（指针类型） 上， 所以 Vertex（值类型） 不满足 `Abser`。
*/

/*
隐式接口
类型通过实现那些方法来实现接口。 没有显式声明的必要；所以也就没有关键字“implements“。

隐式接口解藕了实现接口的包和定义接口的包：互不依赖。

因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。

包 io 定义了 Reader 和 `Writer`；其实不一定要这么做。
*/

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex1{3, 4}
	a = f
	fmt.Println(a.Abs())
	a = &v
	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	//a = v
	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex1 struct {
	X, Y float64
}

func (v *Vertex1) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
