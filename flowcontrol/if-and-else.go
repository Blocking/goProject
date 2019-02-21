package main

import (
	"fmt"
	"math"
)

func pow1(x, y, limit float64) float64 {
	if v := math.Pow(x, y); v < limit {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, limit)
	}
	return limit
}

func main() {
	fmt.Println(
		pow1(3, 2, 10),
		pow1(3, 3, 20),
	)
}
