package main

import (
	"encoding/json"
	"log"
	"fmt"
	"strings"
)

type user struct{
	Name string `json:"name" xml:"name"`
	Age int `json:"age" xml:"age"`
	Phone string `json:"phone" xml:"phone"`
	Password string `json:"-" xml:"-"`
	IsActive bool `json:"active" xml:"active"`
	Role string `json:"-" xml:"-"`
}

var payload = `{"name":"John Smith","age":42,"phone":"","active":true}`

func main(){
	var u user

	enc := json.NewDecoder(strings.NewReader(payload))
	if err := enc.Decode(&u); err != nil{
		log.Fatal(err)
	}
	fmt.Println(u)
}