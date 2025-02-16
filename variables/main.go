
package main

import "fmt"

const FullName string = "Arjun Saseendran" // Public

func main(){
	var username string = "Arjun"
	fmt.Println(username)
	fmt.Printf("The type of username is: %T \n", username)
	
	var active bool = false
	fmt.Println(active)
	fmt.Printf("The type of active is: %T \n", active)

var smallValue uint8 = 255
fmt.Println(smallValue)
fmt.Printf("The type of small value is; %T \n", smallValue)

var smallFloat float32 = 255.725257205720205702
fmt.Println(smallFloat)
fmt.Printf("The type of small float is: %T \n", smallFloat)

var bigFloat float64 = 255.25725025252057702527
fmt.Println(bigFloat)
fmt.Printf("The type of big float is: %T \n", bigFloat)

var anotherVariable int 
fmt.Println(anotherVariable)
fmt.Printf("The value of another variable is:%T \n ", anotherVariable)

var newString string
fmt.Println(newString)
fmt.Printf("The value of new string is: %T \n", newString)

var name = "Arjun"
fmt.Println(name)

age := 34
fmt.Println(age)

fmt.Println(FullName)
fmt.Printf("The full name is: %T\n", FullName)

}


