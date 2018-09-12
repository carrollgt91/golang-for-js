package main

import "fmt"

func main() {
	a := 100

	g := &a // get the memory address for the value in a, store in g
	fmt.Println(g)
	// => 0xc4200140d8
	fmt.Println(*g)
	// => 100
	*g = 200 // this sets the value of a through the pointer
	fmt.Println(a)
	// => 200
}

// See https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back
// for a more thorough and clear explanation.
