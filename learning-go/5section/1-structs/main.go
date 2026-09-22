package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Position  string
	Salary    int
	IsActive  bool
	JoinAt    time.Time
}

func NewEmployee(id int, firstName, lastName, position string, isActive bool) Employee{
	return Employee{
		ID: id,
		FirstName:  firstName,
		LastName: lastName,
		Position: position,
		IsActive: isActive,
		JoinAt: time.Now(),
	}
}

type Contact struct{
	Name string
	Email string
	Phone string
}

func main(){

	// jane := Employee{
	// 	ID: 1,
	// 	FirstName: "Jane",
	// 	LastName: "Doe",
	// 	Position: "Night",
	// 	Salary: 1000,
	// 	IsActive: true,
	// 	JoinAt: time.Now(),
	// }

	// fmt.Printf( "%+v", jane)
	// fmt.Println(jane.ID)
	// fmt.Println(jane.FirstName)
	// fmt.Println(jane.LastName)
	// fmt.Println(jane.Position)
	// fmt.Println(jane.Salary)
	// fmt.Println(jane.IsActive)
	// fmt.Println(jane.JoinAt)

	// joe := NewEmployee(1, "John", "Doe", "Jane", true)

	// fmt.Println(joe.ID)
	// fmt.Println(joe.FirstName)
	// fmt.Println(joe.LastName)
	// fmt.Println(joe.Position)
	// fmt.Println(joe.Salary)
	// fmt.Println(joe.IsActive)
	// fmt.Println(joe.JoinAt)

	// joe.Salary = 1000000
	// fmt.Println(joe.Salary)

	// joePtr := &joe
	// fmt.Println(*joePtr)
	// fmt.Println(joePtr.FirstName)
	// fmt.Println(joePtr.LastName)
	// fmt.Println(joePtr.Position)
	// joePtr.IsActive = true
	// joePtr.LastName = "John Adam"
	// fmt.Println(joe)

	contact1 := Contact{
		Name: "Alice",
		Email: "alice@gmail.com",
		Phone: "111-222",
	}

	contact2 := Contact{
		Name: "Bob",
		Email: "bob@gmail.com",
		Phone: "333-444",
	}

	fmt.Println(contact1)
	fmt.Println(contact2)
	fmt.Println(contact1.Email)
}