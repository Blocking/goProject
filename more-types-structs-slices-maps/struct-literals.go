package main

import "fmt"

type Vertex2 struct {
	X, Y int
}

/*
结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。
*/
var (
	v1 = Vertex2{1, 2}
	v2 = Vertex2{X: 1}
	v3 = Vertex2{}
	p  = &Vertex2{1, 2}
)

func main() {
	m := p.X
	p.X = 2
	n := (*p).X
	(*p).X = 3
	fmt.Println(v1, p, v2, v3, m, n)
}
