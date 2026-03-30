package main

import "fmt"

func main(){
	var revenue float64
	var expense float64
	var taxRate float64
	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expense: ")
	fmt.Scan(&expense)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)
	earningBeforeTax := revenue - expense
	earningAfterTax :=  earningBeforeTax * (1 - taxRate / 100)
	ratio := earningBeforeTax /earningAfterTax
	fmt.Printf("Earning before tax is:%.0f Earning after tax is:%.0f The ratio is: %.0f",earningBeforeTax,earningAfterTax,ratio)
	
}