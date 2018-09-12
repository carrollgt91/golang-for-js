package main

import "fmt"

func main() {
	// this is the syntax for creating a slice - similar to a JS array
	values := []string{"a", "b", "c", "d", "e"}
	values = append(values, "z")

	for index, val := range values {
		fmt.Printf("index %v: %v\n", index, val)
	}

	fmt.Println("first val", values[0])
}
