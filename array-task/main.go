package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {
	hobbies := [3]string{"Coding", "Cricket", "Tuning"}
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)
	courseGoals := []string{"Learn Go", "Learn Rust", "Learn Webdevelopemt with Golang"}
	fmt.Println(courseGoals)
	courseGoals[1] = "Learn Goroutines"
	fmt.Println(courseGoals[1])
	courseGoals = append(courseGoals, "Learn Rust")
	fmt.Println(courseGoals)

	products := []Product{{"1", "iPhone 13", 69999}, {"2", "Macbook", 86000}}
	fmt.Println(products)
	newProduct := Product{
		"23", "Airpods", 14999,
	}
	products = append(products, newProduct)
	fmt.Println(products)

}
