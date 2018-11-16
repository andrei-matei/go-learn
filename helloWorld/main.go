package main

import (
	"fmt"
)

func main() {
	n, err := fmt.Println("This is the true helloWorld")
	fmt.Println("Number of bytes =", n)
	fmt.Println(err)
	foo()
	bar()
}

func foo() {
	fmt.Println("I'm in foo")

	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func bar() {
	x := 42
	fmt.Println("And we exited", x)
}
