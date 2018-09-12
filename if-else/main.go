package main

import "fmt"

func main() {
	// no need for parenthesis unless you're doing more complex
	// boolean expressions
	if b := false; b {
		fmt.Println("You can put an assignment expression in an if block")
	} else {
		fmt.Println("I can still reference b", b)
	} // braces are required, though, even for a single line

	// compiler error - b is out of scope, as variables declared
	// in the first expression of the if statement are only scoped
	// to the if statement
	fmt.Println("But I can't reference it here", b)

	str := ""
	// unlike JS, the expression in an if statement has to resolve to
	// a boolean value
	if len(str) == 0 {
		fmt.Println("It's empty")
	}
}
