package main

import "fmt"

/*
修改 map
在 map m 中插入或修改一个元素：

m[key] = elem
获得元素：

elem = m[key]
删除元素：

delete(m, key)
通过双赋值检测某个键存在：

elem, ok = m[key]
如果 key 在 m 中，`ok` 为 true 。否则， ok 为 `false`，并且 elem 是 map 的元素类型的零值。

同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。
*/
func main() {
	m := make(map[string]int)

	m["Anwser"] = 42
	fmt.Println("The Value:", m["Anwser"])
	m["Anwser"] = 48
	fmt.Println("The Value:", m["Anwser"])
	delete(m, "Anwser")
	fmt.Println("The Value:", m["Anwser"])

	v, ok := m["Anwser"]
	fmt.Println("The value:", v, "Present?", ok)

}
