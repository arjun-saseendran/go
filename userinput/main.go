
package main

import (
"fmt"
"bufio"
"os"
)

func main(){
	welcome := "Welcome to user input!"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")

	input, _ := reader.ReadString('\n')
	fmt.Println("Hello, ",input)
	fmt.Printf("Type of the name is: %T", input)

}
