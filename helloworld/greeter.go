package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Printf("\nHello Golang Plugin\n")
}

var Greeter greeting
