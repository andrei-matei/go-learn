package main

import (
	"fmt"
	"runtime"
)

var x bool

func main() {
	fmt.Println("Lesson6")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = true
	fmt.Println(x)

	a := 7
	b := 42
	fmt.Println(a == b)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
