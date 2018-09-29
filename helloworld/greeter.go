package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Printf("\nHi there welcome to Golang\n")
}

var Greeter greeting
