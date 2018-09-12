package main

import (
	"fmt"
	"net/http"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello, World!</h1>")
}

func main() {
	http.HandleFunc("/hello-world", helloWorld)
	http.ListenAndServe(":3000", nil)
}
