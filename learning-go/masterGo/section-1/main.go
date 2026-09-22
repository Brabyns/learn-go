package main

import "fmt"

type Task struct {
	ID        int
	Completed bool
}

func (t *Task) complete() {
	t.Completed = true
}

//CalculateArea returns the area and a boolean indicationg if inputs were valid

func CalculateArea(length, width float64) (float64, bool) {
	if length <= 0 || width <= 0 {
		return 0, false
	}
	return length * width, true
}

func main() {

	task := Task{ID: 1, Completed: false}

	task.complete()

	fmt.Println(task.Completed)

	area, ok := CalculateArea(10.5, 0)
	if !ok {
		fmt.Println("Invalid dimensions provided")
		return
	}

	fmt.Printf("Area is: %.2f\n", area)
}
