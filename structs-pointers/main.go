package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	p := Person{"Geraldine", 23}
	g := &p // gets the memory address of p and stores it in g

	// Note that we don't use the * operator here - this is a *very*
	// common operation in golang, so rather than forcing you to
	// always explicitly declare that you're setting the value
	// using the * operator each time you make a struct assignment,
	// it simply makes the assignment.
	g.name = "Gerald"

	fmt.Println(p)
	// => {Gerald 23}
}
