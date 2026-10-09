package main

import(
	"fmt"
	"log"
	"encoding/json"
)

type user struct{
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"phone"`
	IsActive bool `json:"active`
}

func main(){
	jane := user{
		Name: "Jane",
		Age: 20,
		IsActive: true,
		Phone: "123-456-789",
	}

	byteSlice, err := json.MarshalIndent(jane, "", " ")
	if err !=nil {
		log.Fatal(err)
	}

	fmt.Println(string(byteSlice))
}