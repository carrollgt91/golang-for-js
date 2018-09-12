package main

import "fmt"

type IntConverter interface {
	Int() int
}

type MyFloat float64

// Note: No "implements" keyword!
type Person struct {
	name string
	age  int
}

func (p *Person) Int() int {
	return p.age
}

func (f MyFloat) Int() int {
	return int(f)
}

func main() {
	var ic IntConverter

	p := Person{"Gerald", 25}

	// causes compiler error:
	// cannot use p (type Person) as type IntConverter in assignment:
	//   Person does not implement IntConverter
	//  (Int method has pointer receiver)
	ic = p

	// this works fine
	ic = &p

	fmt.Println(ic.Int())
	// => 25

	// causes compiler error:
	// ic.name undefined
	//  (type IntConverter has no field or method name)
	fmt.Printf(ic.name)

	// this works fine
	ic = MyFloat(1.5)
}
