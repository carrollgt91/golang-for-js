package main

import "fmt"

func handleMessage(msg interface{}) {
	switch v := msg.(type) {
	case string:
		fmt.Printf("String here %v\n", len(v))
	case int:
		fmt.Printf("Integer here %v\n", v)
	default:
		fmt.Printf("I don't know how to handle this type: %T\n", v)
	}
}

func main() {
	handleMessage("just a string") // => String here just a string
	handleMessage(20)              // => Integer here 20
	handleMessage(false)           // I don't know how to handle this type: bool
}
