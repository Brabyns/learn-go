package main

import "fmt"

func Sum[T int|float32|float64](numbers ...T) T {
	var total T
	for _, number := range numbers{
		total += number
	}
	return total

}

func main() {
	grades := []int{90, 85}
	people := []string{"Jane", "John", "Mark"}

	fmt.Println(len(grades), len(people))

	fmt.Println(Sum(30.3, 34.9))
}