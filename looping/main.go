package main

import "fmt"

func main() {
	// basic for loop
	fact := 1
	for i := 2; i < 8; i++ {
		// notice the lack of colon - we're not redeclaring fact, we're
		// simply assigning it
		fact = fact * i
	}

	fmt.Println("fact", fact)
	// fact 5040

	// what about while?
	// just use for!
	j := 2
	res := 1
	found := false
	for !found {
		res = res * (j)
		j = j + 1
		found = (res == fact)
	}

	fmt.Println("found factorial", res)
	// found factorial 5040

	// infinite loop - just leave it blank
	for {
		// use `break` to escape the loop
	}
}
