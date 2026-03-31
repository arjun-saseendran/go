package main

import "fmt"

func main() {
	anyOutput(20)
	anyOutput("Hello")
	anyOutput(5.5)
	acceptAll("welcome")
	acceptAll(29)
	acceptNew("This is any value")
	checkValue(5)
	result := add(2, 1)
	fmt.Println("The sum is:", result)
}

func anyOutput(value interface{}) {
	switch value.(type) {
	case int:
		fmt.Println("It is integer: ", value)
	case float64:
		fmt.Println("It is float: ", value)
	case string:
		fmt.Println("It is string: ", value)
	}
}

func acceptAll(value interface{}) {
	fmt.Println(value)
}

func acceptNew(value any) {
	fmt.Println(value)
}

func checkValue(value any) {
	intValue, ok := value.(int)
	if ok {
	fmt.Println(intValue + 1, ok)

	}
}

func add[T int | float64 | string ](num1, num2 T)T{
	return num1 + num2
}
