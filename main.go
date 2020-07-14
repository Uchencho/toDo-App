package main

import (
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// Create a function that returns the string below
	fmt.Fprint(w, "Hello Todo App")
}

func getServerAddress() string {
	// Checks if an environment var was passed
	// If not, sets it as :8000

	value, present := os.LookupEnv("PORT")
	if present {
		return value
	}
	return ":8000"
}

func main() {

	http.HandleFunc("/", hello)
	http.ListenAndServe(getServerAddress(), nil)
}
