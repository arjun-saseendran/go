package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name  string `json:"username"`
	Age   int `json:"_"`
	Place string
	Nick  []string `json:"tags,omitempty"`
}

func main() {
	EncodeJson()
}

func EncodeJson() {
	users := []People{
		{"Aswin", 35, "Unknow", []string{"js", "py", "go"}},
		{"Devika", 28, "Unknow", []string{"c", "java", "rust"}},
		{"Maira", 30, "Bridgn", nil},
	}
	// package this data as json
	finalJson, err := json.MarshalIndent(users, "","\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n ",finalJson)
}
