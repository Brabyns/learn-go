package main

import "fmt"

type Contact struct {
	ID    int
	Name  string
	Email string
	Phone string
}

var contactList []Contact
var contactIndexByName map[string]int
var nextID int = 1

func init() {
	contactList = make([]Contact, 0)
	contactIndexByName = make(map[string]int)
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists {
		fmt.Printf("Contact already exists: %v\n", name)
		return
	}

	newContact := Contact{
		ID: nextID,
		Name: name,
		Email: email,
		Phone: phone,
	}

	nextID++
	contactList = append(contactList, newContact)
	contactIndexByName[name] = len(contactList) -1
	fmt.Printf("Contact added: %v\n", name)
}

func findContact(name string) *Contact{
	index, exists := contactIndexByName[name]
	if exists{
		return &contactList[index]
	}
	return nil
}

func ListContact(){
	fmt.Println("--------- List Contacts -----------")
	if len(contactList) == 0{
		fmt.Println("No contact found.")
		return
	}
	for i, contact := range contactList{
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", i+1, contact.ID, contact.Name, contact.Email, contact.Phone)
	}

	fmt.Println("")
}

func main() {
	addContact("Alice Wonderland", "alice@gmail.com", "111-222")
    addContact("Bob Johnson", "bob@gmail.com", "222-333")
    addContact("Charlie Brown", "charlie@gmail.com", "333-444")
    addContact("Diana Prince", "diana@gmail.com", "444-555")
    addContact("Ethan Hunt", "ethan@gmail.com", "555-666")
	addContact("Alice Wonderland", "alice@gmail.com", "111-222")

	ListContact()

	bob := findContact("Bob the Builder")
	if bob == nil {
		fmt.Println("No Bob contact found.")
	}else{
		fmt.Println("Bob contact found.", bob.Name)
	}
}