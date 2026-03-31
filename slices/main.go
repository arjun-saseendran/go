package main

import "fmt"

func main() {
	prices := []float64{1, 2, 3}
	

	newPrices := prices[1:3]
	newPrices[1] = 4
	newPrices = append(newPrices, 5)
	fmt.Println(newPrices, prices)
	prices = append(prices, newPrices...)
	fmt.Println(prices)
	
	users := make([]string, 2, 5)
	users[0]= "Maria"
	users = append(users, "Aswini")
	fmt.Println(cap(users))
	for index, value := range prices{
		fmt.Println("Index is: ",index)
		fmt.Println("Value is: ",value)
	}
}
