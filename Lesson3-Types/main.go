package main

import (
	"fmt"
)

var x = 24
var z = "Shaken not stirred"
var a = "test2"
var i int

func main() {
	y := "test"
	x := "test1"
	i := "integer shit"
	fmt.Println("i=", i)
	//This will print the LOCAL declared x
	fmt.Println("Inside function main: ", x)
	fmt.Println("y =", y)
	fmt.Printf("%T\n", y)
	//fmt.Printf("%T\n", x)
	//fmt.Printf("%T\n", z)
	boo()
}

func boo() {
	//This will print the GLOBAL declared x
	fmt.Println(x)
	fmt.Println("i=", i)
}
