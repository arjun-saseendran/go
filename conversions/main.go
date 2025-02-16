
package main

import  (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Enter user year of birth: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	age, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println("Your age is: ", 2025 - age)
	}

}


