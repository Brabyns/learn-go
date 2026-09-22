package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Contact struct {
	Name  string
	Email string
}

type Post struct {
	ID      int
	Title   string
	Content string
	Author  string
}

func (p Person) Introduce() {
	fmt.Println("Hi, my name is", p.Name)
}

func (p *Person) Birthday() {
	p.Age++
}

func main() {
	person := Person{
		Name: "Alice",
		Age:  25,
	}

	person.Introduce()
	person.Birthday()
	fmt.Println(person.Age)

	contacts := []Contact{
		{Name: "Alice", Email: "alice@gmail.com"},
		{Name: "Bob", Email: "bob@gmail.com"},
		{Name: "Charlie", Email: "charlie@gmail.com"},
	}

	for _, contact := range contacts {
		fmt.Println(contact.Name)
	}

	post := Post{
		ID: 1,
		Title: "Learning Go",
		Content: "Structs are awesome",
		Author: "Brabyns",
	}

	fmt.Println(post.Title)
}
