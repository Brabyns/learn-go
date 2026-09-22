package main

import "fmt"

func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

func add(a, b int){
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
}

func calculateArea(width, height float64) float64{
	if width < 0 || height < 0{
		fmt.Println("Error: width and height must be positive")
		return 0
	}
	return  width * height
}

func calculateArea2(width, height float64) (float64, bool){
	if height < 0 || width < 0 {
		return 0, false
	}

	return width * height, true
}

func factorial(n int) int{
	if n<= 1{
		return 1
	}

	return n* factorial(n-1)
}

func main() {
	greet("Brabyns Yabwetsa")

	add(1, 2)

	area := calculateArea(4, 4)
	fmt.Println(area)

	area2, ok := calculateArea2(10.5, 5.5)

	if !ok {
		fmt.Println("Invalid Demension Provided")
		return
	}
	fmt.Printf("Area is %.2f\n", area2)

	fmt.Println(factorial(5))
}