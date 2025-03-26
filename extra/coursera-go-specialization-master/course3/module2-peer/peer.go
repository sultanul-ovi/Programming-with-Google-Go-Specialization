package main

import (
	"fmt"
	"sync"
)

var counter = 10000

// function to increment a counter
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		counter++
	}
}

// every time this function is called, counter variable will be different since 
func main() {
	wg := sync.WaitGroup{}
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go increment(&wg)
	}
	wg.Wait()
	fmt.Printf("Final counter: %.d\n", counter)

	// datarace occurred 'cause the final counter will be > 10000
}