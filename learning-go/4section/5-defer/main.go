package main

import "fmt"

func simpleDefer() {
	fmt.Println("Fuction simpleDefer: Start")
	defer fmt.Println("Function simpleDefer: defered")
	fmt.Println("Function simpleDefer: Middle")

}

func lifeoSimpleDefer(){
	fmt.Println("FunctionLifoSimpleDefer: Start")
	defer fmt.Println("First Deferred")
	defer fmt.Println("Second: deferred")
	fmt.Println("Fuction lifoSimpleDefer: Middle")
}

func main() {

	defer func(){
		fmt.Println("Before the return of main()")
	}()
	// simpleDefer()
	lifeoSimpleDefer()

	fmt.Println("Last in main()")
}