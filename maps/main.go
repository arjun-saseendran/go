package main

import "fmt"

type mapFloat map[string]float64

func (m mapFloat) output(){
	fmt.Println(m)
}

func main() {
	user := map[string]string{
		"name":  "Maria",
		"place": "Birdgen",
	}
	fmt.Println(user["place"])
	user["age"]= "30"
	delete(user,"place")
	fmt.Println(user)
	
	scores := make(map[string]int32,2)
	scores["maria"] = 10
	scores["aswini"] = 8
	fmt.Println(scores)
	
	marks := make(mapFloat,2)
	marks["aswini"] = 80.5
	marks["devika"] = 70.9
	marks.output()
	for key, value := range user{
		fmt.Println("Key is: ",key)
		fmt.Println("Value is: ",value)
	}
}
