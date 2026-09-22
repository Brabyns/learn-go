package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// filePath :="./10section/1-files/text.txt"
	// data := "Welcome to The GO programming language!"
	// err := os.WriteFile(filePath, []byte(data), 0644)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println("Done Writing")

	content, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}