package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	start := time.Now()
	var t *time.Timer

	// Channel for sync of every goroutine produced with AfterFunc
	flag := make(chan bool)

	t = time.AfterFunc(
		               randomDuration(),
		               func() {
		               fmt.Println(time.Now().Sub(start))
		               flag <- true
		})

	// Infinite loop to reset timer
	for {

		// True only after goroutine is finished
		if <- flag {
			t.Reset(randomDuration())
		}

		// Equivalent of Sleep 5 sec: if more, it finishes
		if time.Since(start) >= time.Duration(5*1e9){
			break
		}
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}

