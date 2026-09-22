package main

import "fmt"

type Person interface{
	GetName() string
}


type BusinessPerson struct {
	ID   int
	Name string
}

func (e BusinessPerson) GetName() string {
	return e.Name
}
func (e BusinessPerson) String() string {
	return fmt.Sprintf("Person[ID:%d, Name:%s]", e.ID, e.Name)
}


func displayPerson(p Person) {
	fmt.Println(p.GetName())
}

type ID int

func (idx ID) String() string{
	return fmt.Sprintf("ID: %d", idx)
}

func main() {
	
	jane := BusinessPerson{
		ID: 1,
		Name: "Jane",
	}

	
	fmt.Println(jane)
	// fmt.Stringer()

	myId := ID(30)
	fmt.Println(myId)
}
