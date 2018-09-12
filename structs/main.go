package main

import "fmt"

// Capital letter means Person is accessible in other packages
type Person struct {
	// lowercase letters mean these data fields are not exposed to other
	// packages
	name string
	age  int
}

func main() {
	p := Person{"Geraldine", 23}
	fmt.Println(p)
	// dot syntax for accessing properties
	fmt.Println(p.name)
	p2 := Person{name: "George"} // age gets default value of 0 here
	fmt.Println(p2)
	p3 := Person{age: 50} // name gets default value of "" here
	fmt.Println(p3)
}
