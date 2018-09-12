package main

import "fmt"

// you can simplify function declarations by
// not repeating types - here, a and b are both
// strings
func split(str string) (string, string) {
	// we'll use the slice operator to split the
	// string into two equal sized halves
	firstHalf := str[:(len(str) / 2)]
	// This slice operator works on many types -
	// strings, slices(similar to JS arrays), and arrays
	secondHalf := str[len(str)/2:]
	return firstHalf, secondHalf
}

func main() {
	s1, s2 := split("YELLOW_SUBMARINE")
	fmt.Println(s1)
	// =>  YELLOW_S
	fmt.Println(s2)
	// =>  UBMARINE
}
