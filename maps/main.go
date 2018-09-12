package main

import "fmt"

func main() {
	lookupTable := map[string]string{
		"key":            "value",
		"differentTypes": "0",
	}

	// This would throw a compiler error the value
	// was a boolean because our map can only store
	// string values
	lookupTable["newVal"] = "true"

	fmt.Println(lookupTable)
	// we can check for whether or not we found a value
	// by using the second return value for a map lookup
	val, ok := lookupTable["key"]
	fmt.Println(val, ok)
	// value true

	delete(lookupTable, "newVal")

	fmt.Println(lookupTable)
}
