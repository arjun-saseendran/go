package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/todos/1"

func main() {

	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("The response type is: %T\n", res)
	defer res.Body.Close()
	databytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println("The response is: ", string(databytes))

}
