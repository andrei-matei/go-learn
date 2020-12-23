package main

import (
	"fmt"

	"github.com/andrei-matei/go-learn/microservice1"

	"github.com/andrei-matei/go-learn/microservice2"
)

func main(){
	fmt.Println("Main module");
	fmt.Println(microservice1.Morning);
	microservice2.Add();
}