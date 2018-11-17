package main

import "fmt"

func ex3() {
	x = 43
	y = "Hello"
	z = true

	fmt.Println("==STARTex3")
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Println("z =", z)
	fmt.Println(x, y, z)
	s := fmt.Sprintf("%T\t %T\t %T\t", x, y, z)
	m := fmt.Sprintf("%v\t %v\t %v\t", x, y, z)
	fmt.Println(m)
	fmt.Println(s)
	fmt.Println("==ENDex3")
}
