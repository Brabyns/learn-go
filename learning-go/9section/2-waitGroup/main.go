package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(message string, delay time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(delay)
	fmt.Println("sayHello", message)
}
func main(){

	var wg sync.WaitGroup

	/*
	1. Add outside of the goroutine
	2. You must decrease the counter by calling wg.Done inside the goroutine
	3. Do not forget to call wg.wait()
	4. Alway pass a refrence of the wait group variable instead of a copy
	*/

	


	wg.Add(4)


	fmt.Println("Hello from main() Goroutine")


	go sayHello("Hello World 1", time.Second, &wg)
	go sayHello("Hello World 2", time.Second, &wg)
	go sayHello("Hello World from  2 seconds", 2*time.Second, &wg)
	go sayHello("Hello World from  3 seconds", 3*time.Second, &wg)


	fmt.Println("Last message from Main() Goroutine")
	wg.Wait()
}