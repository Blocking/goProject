package main

import (
	"fmt"
	"io"
	"os"
)

/**
目前，你已知道 Go 不同于其他编程语言之处。 例如，Go 不需要你在 if、for 或 switch 语句的条件中添加括号。 但是，你始终需要添加大括号 ({})。 可以链接 if 语句，else 子句是可选的。 至关重要的是，可以在 if 条件中声明变量，其作用域仅限 if 块。 即使在同一函数中，也不能访问块外的这些变量。

Go 支持 switch 语句，你不需要编写条件。 只需使用 case 子句即可。 与其他语言不同的是，在 Go 中，无需在每个 case 子句末尾编写 break 语句来避免运行其他 case 子句。

默认情况下，Go 在进入 case 语句后就会运行它，然后退出 switch 子句。 若要跳转到下一个 case 子句，请使用 fallthrough 关键字。 可以从 case 子句调用函数，并且可以在一个 case 子句中将多个表达式分组。

在此模块中，你还了解到，在 Go 中只使用 for 关键字来编写循环。 但是，你可以编写无限循环或 while 条件。 Go 支持 continue 关键字，因此你可以在不退出循环的情况下跳过循环迭代。

最后，你了解了其他 Go 控制流，例如 defer、panic 和 recover 函数。 Go 不支持异常。 它通过使用这三个函数的组合来处理运行时错误。
*/
func courseF() {
	//course1()
	//course2()
	//g(0)
	//fmt.Println("Program finished successfully!")
	/*for i := 1; i <= 100; i++ {
		FizzBuzz(i)
	}*/
	enter()
}

func enter() {
	for {
		val := 0
		fmt.Print("Enter number: ")
		fmt.Scanf("%d", &val)
		switch {
		case val > 0:
			fmt.Println("Yon entered:", val)
		case val == 0:
			fmt.Println("0 is neither negative nor positive")
		case val < 0:
			panic("you got a error")
		}
	}
}

func FizzBuzz(num int) {
	switch {
	case (num%3 == 0) && (num%5 == 0):
		fmt.Println(num, "FizzBuzz")
	case num%3 == 0:
		fmt.Println(num, "Fizz")
	case num%5 == 0:
		fmt.Println(num, "Buzz")
	default:
		fmt.Println(num)
	}
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic("Panic in g() (major)")
	}
	defer fmt.Println("Defer in g()", i)
	fmt.Println("Printing in g()", i)
	g(i + 1)

}

func course2() {
	f, err := os.Create("notes.txt")
	if err != nil {
		return
	}
	defer f.Close()
	if _, err := io.WriteString(f, "Learning GO !"); err != nil {
		return
	}
	f.Sync()
}

func course1() {
	for i := 1; i <= 4; i++ {
		defer fmt.Println("defered", -i)
		fmt.Println("regular", i)
	}
}
