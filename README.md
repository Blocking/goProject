# [_GO_](http://go-tour-zh.appspot.com/list)
## 包、变量和函数。
_学习 Go 程序的基本组件_。
#### 包
 > 程序运行的入口是包 `main`。 倘若不是`main` 程序运行不起来。
 >> 1. 导入包 可以编写多个导入语句 也可以使用`()`导入集合
 >> 2. 在go语言中 首字母大写的名称是被导出的 例如常量`math.Pi`  方法 `fmt.Println`
#### 函数
 > 1. 函数可以接受 多个参数 或者没有参数 变量类型 跟在变量名之后
 > 2. 当俩个或者多个函数命名参数是统一类型 则除了最后一个参数的类型外 其它的参数类型可以省略
 > 3. 方法函数可以返回任意数量的返回值
 > 4. 函数方法的返回值可以当做变量使用 可以被赋值；没有参数的 `return` 语句 直接返回当前 x,y的值 通常不建议这么做 影响代码可读性
#### 变量
 > 1. var 关键字用来定义变量 可以定义在包内以及方法函数内，语法和方法函数的参数列表一致 类型在后面
 > 2. 变量定义可以包含初始值 多个变量初始化时需一一对应;倘若使用表达式进行初始化赋值 则可以省略类型；变量从表达式的初始值中获取类型
 > 3. 可以使用 短声明 符号`:=` 在明确类型的地方 替换 var;但是`:=`不能在函数方法外使用，函数外的每个语句都必须以（`var`、`func`、等等）开始
 > 4. 变量可以打包申明一个语法块儿 如同 "导包"  使用`()`
#### 基本类型
```
      bool
      string
      int  int8  int16  int32  int64
      uint uint8 uint16 uint32 uint64 uintptr
      byte  uint8 的别名
      rune  int32 的别名
            代表一个Unicode码
      float32 float64
      complex64 complex128
      example:
      a := byte(255)  //11111111  0~255正好256个数，255最大
      b := uint8(255) //11111111  0~255正好256个数，不能再高了
      c := int8(127)  //01111111  0~127正好128个数，所以int8的极限只是256的一半 但是int8具有符号 所以范围是-128~127
   ```
 > 2. 倘若不赋值 数值类型的默认值是`0`  布尔类型的默认值为`false` 字符串的默认值为`""`
 > 3. 类型转换 go语言中的类型转换需要显示转换 example: float64(35)
 > 4. go语言存在类型推导语言特性 使用没有类型的 `var` 或 `:=` 语句 由初始值反推变量类型
 > 5. 常量定义使用`const`关键字 来声明 不能使用`:=` ；常量可以是字符、字符串、布尔或者数字类型的值
 > 6. 常量声明定义同样可以使用`()`来完成一个语法块;一个未指定类型的常量由上下文来决定其类型。

## 流程控制语句：for、if、else 和 switch
    学习如何用条件、循环和开关语句控制代码的流程。
 >1 for循环 除了没有`()`其余同java语言一致 `{}`是必须的 例如：

    `	sum := 0
     	for i := 0; i < 10; i++ {
     		sum += i
     	}
     	fmt.Println(sum)`
 >2 跟 C 或者 Java 中一样，可以让前置、后置语句为空。

    sum := 1
    for ;sum<1000 ;  {
      sum += sum;
    }
    fmt.Println(sum)
 >3 基于此可以省略分号：C 的 while 在 Go 中叫做 `for`。 语法基本一致

 >4 如果省略了循环条件，循环就不会结束，因此可以用更简洁地形式表达死循环。 当然也可以在条件初写 `true`

 >5 go语言中的`if` 语句和 `c`以及 `java`中的if基本一致 只不过强制要求不要有`()` 必须有`{}`;
 跟 for 一样，`if` 语句可以在条件之前执行一个简单的语句。
 由这个语句定义的变量的作用域仅在 `if` 范围之内 当然也同时在此 `if` 对应的 `else`范围内。

 >6 `switch`接收的数据类型比`java`多 可以接收布尔类型的数值;`case` 条件匹配可以接收表达式;执行顺序是由上而下的匹配 当匹配成功的时候停止，除非以`fallthrough`关键字结束语句可以使得继续执行下去；
    没有条件的 switch 同 `switch true` 一样。
    这一构造使得可以用更清晰的形式来编写长的 if-then-else 链。

 >7 `defer` 语句会延迟函数的执行直到上层函数返回。 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

## 复杂类型： struct、slice 和 map。
  学习如何基于已有类型定义新的类型：本课涵盖了结构体、数组、slice 和 map。

 >1 `Go`语言中是有指针存在的。指针保存了变量的内存地址：
 `&` 符号会生成一个指向其作用对象的指针;
 `* `符号表示指针指向的底层的值;
 即通常所说的“间接引用”或“非直接引用” 与 C 不同，Go 没有指针运算。

# [TODO](http://go-tour-zh.appspot.com/moretypes/5)






