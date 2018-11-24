package main

import "fmt"

var gogo1 cremvusti
var gogo2 int

func ex5() {
	fmt.Println("==STARTex5")

	fmt.Println(gogo1)
	fmt.Printf("%T\n", gogo1)

	gogo1 = 49
	fmt.Println(gogo1)
	fmt.Printf("%T\n", gogo1)

	gogo2 = int(gogo1)
	fmt.Println(gogo2)
	fmt.Printf("%T\n", gogo2)

	fmt.Println("==ENDex5")
}
