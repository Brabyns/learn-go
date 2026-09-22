package main

import (
	"fmt"
	"sync"
)

func main() {

	counter := 0
	var wg sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i< 10; i++ {
		wg.Add(1)
		go func(){
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}