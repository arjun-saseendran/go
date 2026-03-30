package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	PerformPostFormRequest()
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postform"
	data := url.Values{}
	data.Add("firstname", "Arjun")
	data.Add("lastname", "Saseendran")
	res, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	content, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(content))
}
