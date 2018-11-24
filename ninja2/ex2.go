package main

import (
	"fmt"
)

const (
	x     = 42
	y int = 43
)

func ex2() {
	a := (42 == 42)
	b := (42 <= 43)
	c := (42 >= 43)
	d := (42 != 43)
	e := (42 < 43)
	f := (42 > 43)

	fmt.Println(a, b, c, d, e, f)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	//bit shifting
	z := 79
	fmt.Printf("%d\t%b\t%#X\n", z, z, z)
	q := z << 1
	fmt.Printf("%d\t%b\t%#X\n", q, q, q)

	//raw string literal

	g := "H"
	h := []byte(g)
	fmt.Println(h)
	fmt.Printf("%T\n", h)
}
