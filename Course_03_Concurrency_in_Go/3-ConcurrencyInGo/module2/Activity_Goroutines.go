package main

import (
	"fmt"
	"sync"
)

// Race condition happens when multiple Goroutines try to modify a shared variable without proper synchronization.
// Having a race condition makes results unpredictable.
// To fix a race condition, we can add a Mutex.
// With Mutex, only one Goroutine can modify the shared variable at a time, preventing race condition and guaranteeing a consistent result
var counter int
var wg sync.WaitGroup

// var mu sync.Mutex

func increment() {
	for i := 0; i < 1000; i++ {
		// mu.Lock()
		counter++
		// mu.Unlock()
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	go increment()
	go increment()

	wg.Wait()

	fmt.Printf("Counter: %d\n", counter)
}
