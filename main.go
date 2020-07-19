package main

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
	errorLogger *log.Logger
	infoLogger  *log.Logger
)

func healthcheck(w http.ResponseWriter, req *http.Request) {
	resp := &healthJSON{
		Name:   "REST based TODO APP is up and running",
		Active: true,
	}
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprint(w, string(jsonResp))
}

func theLogger() *os.File {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)

	return file
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

	file := theLogger()

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// defer file.Close()

	http.HandleFunc("/healthcheck", healthcheck)
	infoLogger.Println(getServerAddress())
	if err := http.ListenAndServe(getServerAddress(), nil); err != http.ErrServerClosed {
		errorLogger.Println(err)
	}
}
