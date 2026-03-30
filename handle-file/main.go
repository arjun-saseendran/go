package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Welcome to Golang"
	file, err := os.Create("./golang.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./golang.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside the file is: ", string(databyte))
}
