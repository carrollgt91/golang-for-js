package main

import "fmt"

type Person struct {
	name string
	age  int
}

// Note: we're mixing value receivers and pointer receivers
// here, but this is generally considered an anti-pattern.
func (p Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

func SetName(p *Person, name string) {
	p.name = name
}

func main() {
	p := Person{"Geraldine", 23}
	// notice: we're not using &p to operate on a pointer, even though
	// SetName is a pointer receiver! This is similar to referencing
	// struct fields through pointers - it's a common case, so the
	// language handles it for us
	p.SetName("Gerald")
	// However, we don't get the same convenience when we're using
	// raw functions. This is part of the reason methods are favored
	// over functions in cases where mutating the underlying value
	// is required
	SetName(&p, "Geralde")

	g := &p
	// similarly, we don't need to call *g.GetName() in order for
	fmt.Println(g.GetName())
	// => Geralde
}
