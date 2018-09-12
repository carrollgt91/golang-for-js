package main

import "fmt"

func main() {
	// multiple levels of detail in variable declaration
	// this one is similar to "var" or "let" in JS
	var i int
	i = 1

	// The linter is telling me this is unnecessarily verbose...
	var j int = 2

	// and it is, because this shorthand will use type inference to
	// determine that k is of type int
	k := 3

	fmt.Printf("i %v, j %v, k %v\n", i, j, k)
	// => i 1, j 2, k 3

	// Multiple assignment is an important feature (more on this later)
	a, b := "a", "b"

	fmt.Println(a, b)
	// => a b

	// What happens with uninitialized variables?
	var z int
	fmt.Println("z", z)
	// => z 0
	// int has a default value of 0.

	// The default value for string is the empty string.

	// The default value for bool is false.
	var g bool
	fmt.Println("g", g)
	// => g false
}
