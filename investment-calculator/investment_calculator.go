package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 5.0

func main() {
	var investedAmount float64
	var expectedReturnRate float64 = 6
	var year float64

	fmt.Print("Enter your invested amount: ")
	fmt.Scan(&investedAmount)
	fmt.Print("Enter your year of investment: ")
	fmt.Scan(&year)

	futureValue, futureRealValue := futureValues(inflationRate, investedAmount, expectedReturnRate, year)

	// formatedFutureValue := fmt.Sprintf("%.0f", futureValue)
	// formatedRealValue := fmt.Sprintf("%.0f", futureRealValue)
	// fmt.Println("The future real value is: ", formatedRealValue)
	// fmt.Println("The future value is: ", formatedFutureValue)
	fmt.Println("The future value is: ",futureValue)
	fmt.Println("The real future value is: ",futureRealValue)

}

func futureValues(inflationRate, investedAmount, expectedReturnRate, year float64) (float64, float64) {
	var futureValue = investedAmount * math.Pow(1+expectedReturnRate/100, year)
	var futureRealValue = futureValue / math.Pow(1+inflationRate/100, year)
	return futureValue, futureRealValue
}
