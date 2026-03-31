package main

import "fmt"

func main() {

	prices := [4]float64{
		49999, 69999, 12999, 4999,
	}

	var productNames = [4]string{
		"iPhone", "Mac", "Airpods", "Magic Mouse",
	}
	fmt.Println(productNames)
	specialPrices := prices[1:3]
	discountPrices := specialPrices[:1]
	fmt.Println(discountPrices, specialPrices)
	fmt.Println(len(specialPrices),cap(specialPrices))
	
}
