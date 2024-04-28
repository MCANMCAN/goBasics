package main

import (
	"errors"
	"fmt"
	"time"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

// Below function become method for this struct. It is now part of it.
func (user *user) outputUserDetails() {
	println("First Name is ", user.firstName)
	println("Last Name is ", user.lastName)
	println("Birth Date is ", user.birthDate)
	t := user.createdAt.Format("2006-01-02 15:04:05")

	println("User Created at ", t)

}

// to edit a struct, you need to use pointer .
func (user *user) clearUserDetails() {
	user.firstName = ""
	user.lastName = ""

}

// example of utility-constructer function :
func newUser(firstName, lastName, birthDate string) (*user, error) {
	// adding validation :
	if firstName == "" || lastName == "" {
		return nil, errors.New("firstname and lastname are required")
	}
	return &user{
		firstName: firstName,
		lastName:  lastName,
		birthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func main() {
	userFirstName := getUserData("Please Enter Your First Name: ")
	userLastName := getUserData("Please Enter Your Last Name: ")
	userBirthDate := getUserData("Please Enter Your Birth Date: ")

	var appUser *user
	// appUser = user{
	// 	firstName: userFirstName,
	// 	lastName:  userLastName,
	// 	birthDate: userBirthDate,
	// 	createdAt: time.Now(),
	// }
	appUser, err := newUser(userFirstName, userLastName, userBirthDate)
	if err != nil {
		fmt.Println(err)
		return
	}
	appUser.outputUserDetails()
	appUser.clearUserDetails()
	appUser.outputUserDetails()

}

func getUserData(message string) string {
	var output string
	fmt.Print(message)
	fmt.Scanln(&output)
	return output
}
