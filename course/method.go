package main

import (
	"fmt"
	"geometry"
	"strings"
)

/**
1、方法必须依赖一个结构体
2、调用方法的唯一方式是先声明这个方法所依赖的结构
3、不同结构的方法名 可以重复
4、方法的一个关键点是可以为任何类型定义方法，而不仅仅局限于自定义结构体（如 struct）。但是，你不能通过属于其他包的类型来定义结构。因此不能在
基本类型(如：string)上创建方法。但是我们可以利用一点技巧，基于基本类型来创建自定义类型，然后将其用于基本类型。
*/
func main() {
	t := triangle{3}
	fmt.Println("Perimeter(triangle):", t.perimeter())
	s := square{3}
	fmt.Println("Perimeter(square):", s.perimeter())
	//t.doubleSize()
	//fmt.Println("DoubleSize(triangle):",t.size)
	fmt.Println("Perimeter1(triangle):", t.perimeter1(), t)
	fmt.Println("Perimeter1(triangle):", t)
	u := upperstring("Learning Go!")
	fmt.Println(u)
	fmt.Println(u.Upper())

	c := coloredTriangle{color: "blue", triangle: triangle{3}}
	fmt.Println("Size:", c.size)
	fmt.Println("Perimeter(colored):", c.perimeter())
	//可以显示调用 triangle 中的 perimeter方法
	fmt.Println("Perimeter (normal):", c.triangle.perimeter())
	//方法中的封装 在 Go 中，只需使用大写标识符，即可公开方法，使用非大写的标识符将方法设为私有方法。
	//Go 中的封装仅在程序包之间有效。 换句话说，你只能隐藏来自其他程序包的实现详细信息，而不能隐藏程序包本身。
	m := geometry.Triangle{}
	m.SetSize(3)
	fmt.Println("Perimeter", m.Perimeter())
}

type triangle struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

/**
有时，方法需要更新变量，或者，如果参数过大，则可能尽量避免复制它。遇到此类情况，可以使用指针传递变量的地址。
*/
func (t *triangle) doubleSize() {
	t.size *= 2
}
func (t *triangle) perimeter1() int {
	return t.size * 4
}

type square struct {
	size int
}

func (s square) perimeter() int {
	return s.size * 4
}

/**
不能在基本类型(如：string)上创建方法。但是我们可以利用一点技巧，基于基本类型来创建自定义类型，然后将其用于基本类型。
*/
type upperstring string

func (s upperstring) Upper() string {
	return strings.ToUpper(string(s))
}

/**
嵌入方法
一个结构嵌入另一个结构同时将所有属性嵌入，即可重用结构属性。对于方法也是一致的

实际上是 编译器创建了包装器方法 如：
func (t coloredTriangle) perimeter() int {
    return t.triangle.perimeter()
}
*/
type coloredTriangle struct {
	triangle
	color string
}

/**
重载方法
*/
func (t coloredTriangle) perimeter() int {
	return t.size * 3 * 2
}
