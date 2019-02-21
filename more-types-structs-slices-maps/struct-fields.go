package main

import "fmt"

type Vertex1 struct {
	X int
	Y int
}

func main() {
	v := Vertex1{1, 2}
	v.X = 2
	fmt.Println(v.X)

	p := &v
	p.X = 1e9
	fmt.Println(v)

}
