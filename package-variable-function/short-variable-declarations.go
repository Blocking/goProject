package main

import "fmt"

func main() {
	/**
	可以使用 短声明 符号`:=` 在明确类型的地方 替换 var;但是`:=`不能在函数方法外使用，函数外的每个语句都必须以（`var`、`func`、等等）开始
	*/
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	fmt.Println(i, j, k, c, python, java)
}
