package main

import(
	"os"
	"log"
	"fmt"
)

func main(){
	tempFile, err := os.CreateTemp("", "logs.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func(){
		fmt.Println("Removing tempFile", tempFile.Name())
		_=os.Remove(tempFile.Name())
	}()

	_, err = tempFile.Write([]byte("Hello World\n"))
	if err != nil {
		log.Fatal(err)
		tempFile.Close()
		return
	}

	// TEMP DIR

	tempDir, err := os.MkdirTemp("", "My_app_logs")
	if err !=nil{
		log.Fatal(err)
	}

	defer func(){
		fmt.Println("Removing tempDir", tempDir)
		_=os.Remove(tempDir)
	}()
}