package main

import (
	"fmt"
	"log"
	"os"
	"io"
	"bufio"
)

func main() {

	// ========================Writing to a file named "text.txt"========================


	filePath :="./text.txt"
	data := "Welcome to The GO programming language!"
	err := os.WriteFile(filePath, []byte(data), 0644)
	if err != nil {
		log.Fatal(err)
	}


	// ========================Reading from a file named "text.txt"========================

	fmt.Println("Done Writing")

	content, err := os.ReadFile("text.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))


	// ========================Creating and writing to a new file named "text2.txt"========================

	/*
	file2, er := os.Create("file-via-create.txt")
	if er != nil {
		log.Fatal(er)
	}
	defer file2.Close()

	_, err = file2.WriteString("This is some content for file-via-create.txt")
	if err != nil {
		log.Fatal(err)
	}

	*/


	// ===============================CREATE A LIST ===============================
	/*
	newFileContent, err := os.Open("file-via-create.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer newFileContent.Close()

	scanner := bufio.NewScanner(newFileContent)
	lineNum := 1
	for scanner.Scan() {
		lineNum++
		fmt.Println(lineNum, scanner.Text())
	}

	if err := scanner.Err(); err != nil{
		if err != io.EOF{
			log.Fatal(err)
		}
	}
		*/



	//=========================APPEND NEW DATA ==============================

	fileName := "file-via-create.txt"

	printContent(fileName)

	newFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil{
		log.Fatal(err)
	}
	defer newFile.Close()

	_,_ = newFile.WriteString(fmt.Sprintf("- C\n"))
	_,_ = newFile.WriteString(fmt.Sprintf("- Ruby\n"))
	_,_ = newFile.WriteString(fmt.Sprintf("- Fortan\n"))
	_,_ = newFile.WriteString(fmt.Sprintf("- Ada\n"))
	_,_ = newFile.WriteString(fmt.Sprintf("- Rust\n"))


}

func printContent(fileName string){
	newFile, err := os.Open(fileName)
	if err != nil{
		log.Fatal(err)
	}
	defer newFile.Close()

	scanner := bufio.NewScanner(newFile)
	lineNum := 1
	for scanner.Scan(){
		fmt.Println(lineNum, scanner.Text())
		lineNum++
	}

	if err := scanner.Err(); err !=nil{
		if err != io.EOF{
			log.Fatal(err)
		}
	}
}