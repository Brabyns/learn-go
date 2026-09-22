package main 

import "fmt"

func main(){
	tmp := 25
	if tmp > 30{
		fmt.Println("Greater than 30")
	}else{
		fmt.Println("Lesser than 30")
	}

	userAccess := map[string]bool{
		"jane": true,
		"john": false,
	}

	if hasAccess, ok := userAccess["jane"]; ok && hasAccess {
		fmt.Println("Jane can access the system")
	}else{
		fmt.Println("access not granted")
	}
}