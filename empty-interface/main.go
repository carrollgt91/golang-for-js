package main

import (
	"fmt"
)

func main() {
	var a interface{}

	a = "A string... for now"
	fmt.Println(a)

	a = 200
	// Notice println accepts both values of a...
	// wonder what it's signature is?
	fmt.Println(a)
}
