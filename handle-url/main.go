package main

import (
	"fmt"
	"net/url"
)

func main() {
	const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"
	// parsing
	result, err := url.Parse(myurl)
	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
	qparams := result.Query()
	fmt.Printf("The type of the query params are: %T\n", qparams)
	fmt.Println(qparams["coursename"])
	for val := range qparams {
		fmt.Println("Params is: ", val)
	}
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
