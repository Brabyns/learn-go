package main

import "fmt"

func mightPanic(shouldPanic bool) {
	if shouldPanic {
		panic("Something went wrong in mightPanic")
	}
	fmt.Println("This function executed without a panic")
}

func recoverable(){
	defer func(){
		if r := recover(); r !=nil {
			fmt.Println("Recovered from panic:",r)
		}
	}()

	mightPanic(true)
}

func main() {
	// panic("Something bad happened")
	// mightPanic(false)
	recoverable()
}