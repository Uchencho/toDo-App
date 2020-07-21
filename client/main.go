package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const url = "http://127.0.0.1:8000/healthcheck"
const expectedBody = `{"Name":"REST based TODO APP is up and running","Active":true}`

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
	if resp.StatusCode != 200 || string(body) != expectedBody {
		log.Fatalln(resp.StatusCode, string(body))
	}

	fmt.Printf("Status:	%d\n", resp.StatusCode)
	fmt.Printf("Body: %s\n", body)
}
