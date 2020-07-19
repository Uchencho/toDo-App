package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://127.0.0.1:8000/healthcheck"

func main() {
	resp, err := http.Get(url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	fmt.Printf("Status:	%d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)
}
