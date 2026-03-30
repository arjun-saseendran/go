package main

import (
	"fmt"
	"github.com/arjunsaseendran/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birth date (MM/DD/YYYY): ")

	


	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Print(err)
		return
	}
	
	admin := user.NewAdmin("go@dev.com","1234")
	admin.ClearName()

	// appUser.OutputUserDetails()
	// appUser.ClearName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
