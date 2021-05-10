package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += 1
	}
	fmt.Println(sum)

	var num int64
	rand.Seed(time.Now().Unix())
	for num != 5 {
		num = rand.Int63n(15)
		fmt.Println(num)
	}

	var index int32
	for {
		fmt.Println("Writting inside the loop")
		if index = rand.Int31n(10); index == 5 {
			fmt.Println("finish!")
			break
		}
		fmt.Println(index)
	}
	contineExercise()
}

func contineExercise() {
	sum := 0
	for num := 1; num <= 100; num++ {
		if num%5 == 0 {
			continue
		}
		sum += num
	}
	fmt.Println("The sum of 1 to 100 , but excluding numbers divisible by 5, is ", sum)

}
