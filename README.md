# [_GO_](http://go-tour-zh.appspot.com/list)
## 包、变量和函数。
_学习 Go 程序的基本组件_。
#### 包
 > 程序运行的入口是包 `main`。 倘若不是`main` 程序运行不起来。
 >> 1. 导入包 可以编写多个导入语句 也可以使用`()`导入集合
 >> 2. 在go语言中 首字母大写的名称是被导出的 例如常量`math.Pi`  方法 `fmt.Println`  
 >> 3. 你现在可以开始编写包的函数和变量。 不同于其他编程语言，Go 不会提供 public 或 private 关键字，以指示是否可以从包的内外部调用变量或函数。 但 Go 须遵循以下两个简单规则：
 >> + 如需将某些内容设为专用内容，请以小写字母开始。
 >> + 如需将某些内容设为公共内容，请以大写字母开始。
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

    	sum := 0
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
   `fallthrough`会使得switch立马转到下一个case语句 **_而不会对case条件进行验证_**

 >7 `defer` 语句会延迟函数的执行直到上层函数返回。 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。
 延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

## 复杂类型： struct、slice 和 map。
  学习如何基于已有类型定义新的类型：本课涵盖了结构体、数组、slice 和 map。

 >1 `Go`语言中是有指针存在的。指针保存了变量的内存地址：  
 `&` 符号会生成一个指向其作用对象的指针;  
 `* `符号表示指针指向的底层的值;
即通常所说的“间接引用”或“非直接引用” 与 C 不同，Go 没有指针运算。
 + `&` 运算符使用其后对象的地址。
 + `*` 运算符取消引用指针。 也就是说，你可以前往指针中包含的地址访问其中的对象。

# [TODO](http://go-tour-zh.appspot.com/moretypes/5)


# 微软课程
## 练习 - 使用映射
大体来说，Go中的映射是一个哈希表，是键值对的集合。映射中过所有的键都必须具有相同的类型，它们的值也是如此。
不过，可对键和值使用不同的类型。例如，键可以是数字，值可以是字符串。若要访问映射中的特定项，可引用该项的键。
## 声明和初始化映射
若要声明映射，需要使用`map`关键字。然后，定义键和值类型，如下所示：`map[T]T`。例如，如果要创建一个包含学生年龄的
映射，可使用以下代码：

## 总结
你在本模块中学习了 4 种数据类型，它们将帮助你更好地表示你的程序将用于解决问题的数据。 我们首先了解了 Go 中的数组，如你所见，它很简单。 如果你一直在用其他语言编程，也没有太大区别。 但你需要了解数组的工作原理，才能理解我们探索的其他 3 种数据类型。
例如，你已了解到切片是一个简单的数据结构，它有一个指针指向一个基础数组，有两个属性用于控制该数组的长度和容量。 你不必担心切片的大小，因为 Go 会帮你“扩展”基础数组的大小。 你还了解到切片运算符可帮助你创建新的子切片并从切片中删除元素。
接着我们了解了映射，这是一种类似于切片和数组的数据结构。 区别是映射由键或值元素组成，其中键和值可以是不同的类型。 你还有一个内置函数可用于快速从映射中删除元素。 如果你尝试从不存在的映射访问位置，Go 不会引发 panic 错误。
最后，我们了解了 Go 中的结构，我们将在下一模块中继续讲解这一知识。 结构是不同类型的字段的集合，这些字段可用于表示数据库中的项。 如果需要，你也可将结构转换为 JSON 格式。
在后续模块中，我们将继续使用这些数据类型来探索 Go 中的其他功能。

# 在 Go 中实现错误处理和日志记录
### 用于错误处理的推荐做法
在 Go 中处理错误时，请记住下面一些推荐做法：
+ 始终检查是否存在错误，即使预期不存在。 然后正确处理它们，以免向最终用户公开不必要的信息。
+ 在错误消息中包含一个前缀，以便了解错误的来源。 例如，可以包含包和函数的名称。
+ 创建尽可能多的可重用错误变量。
+ 了解使用返回错误和 panic 之间的差异。 不能执行其他操作时再使用 panic。 例如，如果某个依赖项未准备就绪，则程序运行无意义（除非你想要运行默认行为）。
+ 在记录错误时记录尽可能多的详细信息（我们将在下一部分介绍记录方法），并打印出最终用户能够理解的错误。

## 了解如何在 Go 中记录

### 总结

如你所见，Go 中的错误处理和日志记录与其他编程语言中的这些过程不同。 首先，Go 的错误处理方法非常简单。 使用 if 条件，调用的函数应返回多个值。 按照惯例，最后一个返回值为错误。 如果错误变量返回 nil，则不存在错误。 如果值不为 nil，则存在失败。 只需再次返回错误即可将错误传播到堆栈，并且可以根据需要添加更多上下文。
可以创建可重用为程序中常见错误消息的返回值的错误变量。
你还需要了解何时使用 panic。 我们已介绍 panic 和 recover 的工作原理。 仅当明确需要停止程序时，才应使用这些函数。 有时，即使你正确处理了错误，程序也可能会停止响应。 但这应该是异常，而不是规则。
最后，我们探讨了 Go 中日志记录的工作原理，你了解了如何使用标准库。 除了将日志打印到控制台之外，你还可以将日志发送到文件供稍后处理，然后将它们发送到一个集中位置。 当代码库扩大时，你可能需要执行其他操作，例如设置日志级别或配置不同输出。 标准库中不支持这些任务。 你将需要使用记录框架，例如 zerolog。
此模块较短，但请务必充分了解相关概念。 在需要对程序中的问题进行故障排除时，它们将有所帮助。

# 在 Go 中使用方法
 GO 中的方法是一种特殊的函数 与函数的区别是 你必须在函数名称之前加入一个额外的参数。此参数称为接收方。

## 声明方法
语法：

    func (variable type) MethodName(parameters ...) {
       // method functionality 
    }
+ 声明方法之前 需要先创建方法接收的结构体 即 方法是挂在结构体下的
+ 方法必须依赖一个结构体 
+ 调用方法的唯一方式是先声明这个方法所依赖的结构 
+ 不同结构的方法名 可以重复
+ 方法的一个关键点是可以为任何类型定义方法，而不仅仅局限于自定义结构体（如 struct）。但是，你不能通过属于其他包的类型来定义结构。因此不能在
  基本类型(如：string)上创建方法。但是我们可以利用一点技巧，基于基本类型来创建自定义类型，然后将其用于基本类型。
+ 有时，方法需要更新变量，或者，如果参数过大，则可能尽量避免复制它。遇到此类情况，可以使用指针传递变量的地址。
+ 不能在基本类型(如：string)上创建方法。但是我们可以利用一点技巧，基于基本类型来创建自定义类型，然后将其用于基本类型。
+ 嵌入结构体 不止属性嵌入进去 同时 以编译器创建了包装器方法的形式嵌入进去 表现在外在就是被继承了 其实是编译器在子结构体中创建了包装器方法 将嵌入结构体的方法进行了包装
，当然你可以在子结构体中写一个同样名称的方法 这样表现在外在就是该方法被子结构重写了。实际上是俩个结构各自有各自的方法 它们之间的联系是由编译器生成包装器方法进行串联的。
  
### 方法中的封装 
 在GO中只需使用大写标识 即可公开方法 使用非大写标识将方法设为私有    
Go中的封装仅在程序包之间有效。

# 在 Go 中使用接口
接口是表示其它类型行为的数据类型。接口是一个约定 使得代码结构更加灵活。 同时在GO中接口是满足隐式实现的，不需要显示的特别指定。
## 声明接口
使用 关键字 `interface`
## 实现接口
+ 当类型的方法满足接口当中的所有方法时，则满足按类型的隐式实现。
+ 接口可以多实现。
+ 接口可用于自定义扩展


# 并发
## 了解 goroutine（轻量线程）





# 微软学习课程 -> [masteringGo](https://github.com/hantmac/Mastering_Go_Second_Edition_Zh_CN) -> 极客时间go课程






