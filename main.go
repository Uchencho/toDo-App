package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// Create a function that returns the string below
	fmt.Fprint(w, "Hello Todo App")
}

func main() {
	// Registering a handler
	http.HandleFunc("/", hello)

	// Since a handler has been registered above, we can pass
	// nil as the second arguement to listen and server
	http.ListenAndServe(":8000", nil)
}
