package main

import "fmt"

// you can simplify function declarations by
// not repeating types - here, a and b are both
// strings
func concat(a, b string) string {
	return a + " " + b
}

func main() {
	res := concat("Hello", "World")
	fmt.Println(res)
	// => Hello World
}
