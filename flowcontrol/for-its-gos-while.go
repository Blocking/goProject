package main

import "fmt"

func main() {
	sum := 1
	index := 0
	for sum < 1000 {
		sum += sum
		index = index + 1
	}
	fmt.Println(sum)
	fmt.Println(index)
}
