package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, req *http.Request) {
	// Create a function that returns the string below
	fmt.Fprint(w, "Hello Todo App")
}

type healthJSON struct {
	Name   string
	Active bool
}

func healthcheck(w http.ResponseWriter, req *http.Request) {
	resp := &healthJSON{
		Name:   "REST based TODO APP is up and running",
		Active: true,
	}
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprint(w, string(jsonResp))
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
