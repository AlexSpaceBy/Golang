package main

import (
	"fmt"
	"sync"
)

var sl []string
var wg sync.WaitGroup

// To get Mutex.Lock() Mutex.Unlock() methods
var mu sync.Mutex

func addLine(words string) {

	// Lock shared memory
	mu.Lock()

	sl = append(sl, words)

	//Unlock shared memory
	mu.Unlock()
	wg.Done()
}

func main() {
	wg.Add(4)
	go addLine("I'll")
	go addLine(" be here")
	go addLine(" all day")
	go addLine(" and you'll be too")
	wg.Wait()
	fmt.Println(sl)
}
