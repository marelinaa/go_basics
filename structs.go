package main

import (
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthdate string
	createdAt time.Time
}

// (u user) is called receiver argument
func (u *user) outputUserDetails() {
	//technically correct way to use data not the adress of u
	//fmt.Println((*u).firstName, (*u).lastName, (*u).birthdate)
	//next way is like shortcut, but allowed by Go
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *user) clearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser user = user{
		firstName: userFirstName,
		lastName:  userLastName,
		birthdate: userBirthdate,
		createdAt: time.Now(),
	}

	// ... do something awesome with that gathered data!

	appUser.outputUserDetails()
	appUser.clearUserName()
	appUser.outputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
