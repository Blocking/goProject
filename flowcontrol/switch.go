package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("GO runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux:":
		fmt.Println("Linux.")
	case "windows":
		fmt.Println("windows11")
		//fallthrough
	default:
		fmt.Printf("%s.", os)
	}
}
