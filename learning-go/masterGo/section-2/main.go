package main

import "fmt"

type Users struct {
	Name  string
	Email string
}

func (U *Users) UpdateEmail(newEmail string) {
	U.Email = newEmail
}

func main() {
	user := Users{
		Name:  "Jane Doe",
		Email: "jane@gmail.com",
	}

	fmt.Printf("Before: %+v\n", user)

	user.UpdateEmail("john@gmail.com")
	fmt.Printf("After: %+v\n", user)
}
