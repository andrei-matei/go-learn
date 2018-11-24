package main

import "fmt"

type cremvusti int

var gogo cremvusti

func ex4() {
	fmt.Println("==STARTex4")

	fmt.Println(gogo)
	fmt.Printf("%T\n", gogo)

	gogo = 45
	fmt.Println(gogo)
	fmt.Printf("%T\n", gogo)

	fmt.Println("==ENDex4")
}
