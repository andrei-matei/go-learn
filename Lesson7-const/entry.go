package main

import "fmt"

func main() {
	const (
		a = 72
		b = 72.38
		c = "James Bond"
		d = true
	)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

}
