package main

import (
	"fmt"
	
)

func main() {
	day := "Thursday"
	fmt.Println("Today is ", day)

	switch day {
	case "Suday", "Saturday":
		fmt.Println("Weekend! No work")

	case "Moday", "Tuesday":
		fmt.Println("Work days. Lots of meetings")
	default:
		fmt.Println("Mid-week")
	}

	checkType := func(i interface{}){
		switch v := i.(type){
		case int:
			fmt.Printf("Integer: %d\n", v)
		case string:
			fmt.Printf("String: %s\n", v)
		case bool:
			fmt.Printf("Boolean: %t\n", v)
		default:
			fmt.Printf("Unknown type: %T\n", v)
		}
	}

	checkType(22)
	checkType("Test")
	checkType(true)
	checkType(314.47)

}