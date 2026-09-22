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
//FullName A value receiver
func (e Employee) FullName() string {
	return e.FirstName + " " + e.LastName
}

func (e *Employee) Deactivate()  {
	e.IsActive = false
}

func deactivate(e *Employee){
	e.IsActive = false
}

func (e *Employee) SetJoinedAt ( t time.Time){
	e.JoinAt = t
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

func main(){

	jane := Employee{
		ID: 1,
		FirstName: "Jane",
		LastName: "Doe",
		Position: "Night",
		Salary: 1000,
		IsActive: true,
		// JoinAt: time.Now(),
	}

	fmt.Println(jane.FullName())

	fmt.Printf("%+v\n", jane)
	jane.Deactivate()
	fmt.Printf("%+v\n", jane)
	// deactivate(&jane)
	// fmt.Printf("%+v\n", jane)

	jane.SetJoinedAt((time.Now().Add(10000000 * time.Minute)))
	fmt.Printf("%+v\n", jane)

	joe := &Employee{}
	joe = nil  // Demo to show panic

	joe.FullName() //Bad idea



}