package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const expectedBody = `{"Name":"REST based TODO APP is up and running","Active":true}`

func main() {

	// Build first then call ./main -baseurl=http://127.0.0.1:8888
	baseurl := flag.String("baseurl", "http://127.0.0.1:8000", "baseurl of resource")
	flag.Parse()
	url := *baseurl + "/healthcheck"

	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error in getting resource: ", err)
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
