package main

import "github.com/carrollgt91/golang-for-js/packages/mypackage"

func main() {
	// compiler errors on this line
	// cannot refer to unexported name mypackage.private
	mypackage.private()
	// This works fine
	mypackage.Public()
}
