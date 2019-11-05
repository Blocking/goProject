package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

/*
这么写会陷入死循环。
仔细想了想，应该是无限递归陷入的死循环
因为e变量是一个通过实现Error()的接口函数成为了error类型，那么在fmt.Sprint(e)时就会调用e.Error()来输出错误的字符串信息
于是函数相当于
func (e ErrNegativeSqrt) Error() string {
return fmt.Sprint(e.Error())
}
从而陷入了无限递归之中

作者：lucasdada
链接：https://www.jianshu.com/p/c77730e03393
来源：简书
简书著作权归作者所有，任何形式的转载都请联系作者获得授权并注明出处。
*/

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("Cannot Sqrt negative number:", e)
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	} else {
		return math.Sqrt(x), nil
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
