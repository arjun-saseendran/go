package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000"
	res, err := http.Get(myUrl)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	fmt.Println("The status code of the response: ", res.StatusCode)
	fmt.Println("The content length is: ", res.ContentLength)
	var responseString strings.Builder
	content, err := ioutil.ReadAll(res.Body)
	bytecount, _ := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	// fmt.Println("The data is: ",string(content))
	fmt.Println("The data is: ", bytecount)
	fmt.Println(responseString.String())
}
