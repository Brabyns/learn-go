package main

import (
	"fmt"
	"strings"
)

func main(){
	s1 := "abc"
	s2 := strings.Clone(s1)

	fmt.Println(s1, s2)

	parts := strings.Split("test@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println(username, domain)


	parts = strings.Fields("jane example.com")
	username, domain = parts[0], parts[1]
	fmt.Println(username, domain)
}