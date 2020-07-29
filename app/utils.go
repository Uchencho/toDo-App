package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type healthJSON struct {
	Name   string
	Active bool
}

var (
	ErrorLogger *log.Logger
	InfoLogger  *log.Logger
)

func Healthcheck(w http.ResponseWriter, req *http.Request) {
	resp := &healthJSON{
		Name:   "REST based TODO APP is up and running",
		Active: true,
	}
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprint(w, string(jsonResp))
}

func TheLogger() *os.File {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	return file
}

func GetServerAddress() string {

	const defaultServerAddress = "127.0.0.1:8000"
	serverAddress, present := os.LookupEnv("SERVER_ADDRESS")
	if present {
		return serverAddress
	}
	return defaultServerAddress
}
