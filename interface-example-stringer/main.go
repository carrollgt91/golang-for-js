package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v : %v years old.", p.name, p.age)
}

func main() {
	g := Person{"Gabby", 72}
	r := Person{"Ronald", 93}
	fmt.Println(g, r)
	// Gabby: 72 years old. Ronald: 93 years old.
}
