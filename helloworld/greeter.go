package main

import "fmt"

type greeting string

func (g greeting) Greet() {
	fmt.Printf("\nHi there welcome to Golang\n")
}

func cGreet() {
	fmt.Printf("\nThis golang method can call from C, Python or another languages\n")
}

var Greeter greeting

func main() {}
