package main

import (
	"fmt"
)

type hotdog int

var cremvusti hotdog = 42
var carnat hotdog = 43

func main() {
	fmt.Println(cremvusti)
	fmt.Printf("%T\n", cremvusti)

	fmt.Println(carnat)
	fmt.Printf("%T\n", carnat)

}
