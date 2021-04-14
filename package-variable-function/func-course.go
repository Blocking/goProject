package main

/**
将值传递给函数时，该函数中的每个更改都不会影响调用方。
Go 是“按值传递”编程语言。 这意味着每次向函数传递值时，Go 都会使用该值并创建本地副本（内存中的新变量）。
在函数中对该变量所做的更改都不会影响你向函数发送的更改。
*/
func main() {
	firstName := "John"
	updateName(firstName)
	println(firstName)
	updateNamePointer(&firstName)
	//println(firstName)
	println(&firstName)
}

/**
& 运算符使用其后对象的地址。
* 运算符取消引用指针。 也就是说，你可以前往指针中包含的地址访问其中的对象。
*/
func updateNamePointer(s *string) {
	println(s)
	println(*s)
	*s = "David"

}

func updateName(name string) {
	name = "David"
}
