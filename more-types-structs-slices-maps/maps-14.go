package main

/*
map 映射键到值。

map 在使用之前必须用 make 而不是 new 来创建；值为 nil 的 map 是空的，并且不能赋值。
*/
import "fmt"

type Vertex5 struct {
	Lat, Long float64
}

var m map[string]Vertex5

/*
	map 的文法跟结构体文法相似，不过必须有键名。
*/
var n = map[string]Vertex5{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	m = make(map[string]Vertex5)
	m["Bell Labs"] = Vertex5{40.68433, -74.39967}
	fmt.Println(m["Bell Labs"])
	fmt.Println(n)

}
