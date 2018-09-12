package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) SetName(name string) {
	p.name = name
}

func (p Person) GetName() string {
	return p.name
}

func GetName(p Person) string {
	return p.name
}

func main() {
	p := Person{"Geraldine", 23}
	fmt.Printf(p.GetName())
	// => Geraldine
	p.SetName("Gerald")
	fmt.Printf(p.GetName())
	// => Gerald
	fmt.Printf(GetName(p))
	// => Gerald
}
