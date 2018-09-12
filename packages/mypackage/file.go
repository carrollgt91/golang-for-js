package mypackage

import "fmt"

func private() {
	fmt.Println("Can't be called outside this package")
}

// Public is a trivial example of a public function.
// My IDE told me I needed to document this function, so I did.
func Public() {
	fmt.Println("Available for the world to see")
}
