package main

import (
	"fmt"
	"sync"
)

// To prevent race condition
var mu sync.Mutex

func main() {

	// Go sync method
	var wg sync.WaitGroup


	incrementer := 0
	gs := 100

	// We have 100 goroutines to sync
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {

			// To prevent race conditions
			mu.Lock()
			defer mu.Unlock()

			v := incrementer
			v++
			incrementer = v
			fmt.Println(incrementer)

			// Goroutine is done flag
			wg.Done()
		}()
	}

	// When all goroutine are done the program continue
	wg.Wait()
	fmt.Println("end value:", incrementer)
}
