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
	// Checks if a server address was passed as an
	// environment variable
	// If not, sets it as "127.0.0.1:8000"

	value, present := os.LookupEnv("SERVER_ADDRESS")
	if present {
		return value
	}
	return "127.0.0.1:8000"
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(getServerAddress(), nil)
}
