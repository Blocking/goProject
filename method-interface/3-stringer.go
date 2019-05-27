package main

import (
	"fmt"
	"strconv"
)

/*
Stringers
一个普遍存在的接口是 fmt 包中定义的 Stringer。

type Stringer interface {
    String() string
}
Stringer 是一个可以用字符串描述自己的类型。`fmt`包 （还有许多其他包）使用这个来进行输出。
*/
type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	addrs := map[string]IpAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for n, v := range addrs {
		fmt.Printf("%v:%v\n", n, v)
	}

}

type IpAddr [4]byte

func (ip IpAddr) String() string {
	var s string
	for i := 0; i < len(ip); i++ {
		if i == 0 {
			s = strconv.Itoa(int(ip[i]))
		} else {
			s = s + "." + strconv.Itoa(int(ip[i]))
		}
	}
	return s
}
