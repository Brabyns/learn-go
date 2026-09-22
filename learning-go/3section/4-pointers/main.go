package main

import "fmt"

func modifyValue(val int){
	val = val *10
	fmt.Printf("modifyValue: %+v\n", val)
}

func modifyPointer(val *int){
	if val == nil{
		fmt.Println("Val is nil")
		return
	}

	*val = *val * 10 //Dereference
	fmt.Printf("modifyPointer: %+v\n", *val)
}

func main() {
	// age := 10
	// ageptr := &age
	// fmt.Printf("age: %d\n", &age)
	// fmt.Printf("ageptr: %d\n", ageptr)

	num := 10
	modifyValue(num)
	fmt.Println(num)
	modifyPointer(&num)
	fmt.Println(num)
}