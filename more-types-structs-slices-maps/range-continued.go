package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		b := uint(i)
		c := 1 << b
		fmt.Println(i, b, c)
		pow[i] = c
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
