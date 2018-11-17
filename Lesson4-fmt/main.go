package main

import (
	"fmt"
)

var y = 42

func main() {
	//Print, Printf, Println
	fmt.Println(y)
	fmt.Printf("%d\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)

	//Sprint, Sprintf, Sprintln
	y := 911
	s := fmt.Sprintf("%d\t %b\t %x\n", y, y, y)
	fmt.Println(s)

}
