package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name  string   `json:"name"`
	Age   int      `json: "age"`
	Place string   `json: "place"`
	Tags  []string `json: "tags"`
}

func main() {
DecodeJson()
}

func DecodeJson() {
	jsonData := []byte(`
		{
		"name":"Arjun",
		"age": 34,
		"place": "Cherthala",
		"tags": ["js", "go", "rust", "py"]
		}

		`)
	var users People
	checkJsonValid := json.Valid(jsonData)
	if checkJsonValid {
		fmt.Println("Json is valid!")
		json.Unmarshal(jsonData, &users)
		fmt.Printf("%#v\n", users)
	}else{
		fmt.Println("Json not valid!")
	}
	
	// Handle using maps
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonData,&myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)
	for k,v := range myOnlineData {
		fmt.Printf("The key is %v and the value is %v also type is %T\n", k,v,v)
	}
	
}
