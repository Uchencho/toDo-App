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

func healthcheck(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "REST based TODO APP is up and running")
}

func getServerAddress() string {

	const defaultServerAddress = "127.0.0.1:8000"
	serverAddress, present := os.LookupEnv("SERVER_ADDRESS")
	if present {
		return serverAddress
	}
	return defaultServerAddress
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/healthcheck", healthcheck)
	http.ListenAndServe(getServerAddress(), nil)
}
