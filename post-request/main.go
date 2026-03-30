package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformPostRequest()
}

func PerformPostRequest() {
	const myurl = "http://localhost:3000"
	// fake json payload
	requestBody := strings.NewReader(`
	{	"name": "Arjun",
		"age": 34
	}	`)
	res, err := http.Post(myurl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
