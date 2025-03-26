package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	philosopherCount = 5
	eatTimes         = 3
	maxEating        = 2
)

type Philosopher struct {
	number   int
	left     *sync.Mutex
	right    *sync.Mutex
	host     chan struct{}
	eatCount int
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eatCount < eatTimes {				
		// wait for permision token from the host
		token := <-p.host
		
		// Pick up chopsticks in any order
		p.left.Lock()		
		p.right.Lock()		

		// Start eating
		fmt.Printf("starting to eat %d\n", p.number)
		time.Sleep(time.Millisecond * 100) // Simulate eating time

		// Finish eating
		fmt.Printf("finishing eating %d\n", p.number)
		p.eatCount++

		// Put down chopsticks
		p.right.Unlock()
		p.left.Unlock()

		// Return the permission token to host
		p.host <- token
	}
}

func main() {
	// Create chopsticks
	chopsticks := make([]sync.Mutex, philosopherCount)

	// Create a host buffered channel to allow only 2 philosophers to eat concurrently
	host := make(chan struct{}, maxEating)

	// Create philosophers
	philosophers := make([]*Philosopher, philosopherCount)

	// Assing chopsticks to  philosophers
	for i := 0; i < philosopherCount; i++ {
		philosophers[i] = &Philosopher{
			number:   i + 1,
			right:    &chopsticks[i],
			left:     &chopsticks[(i+1)%philosopherCount],
			host:     host,
			eatCount: 0}
	}

	// Send 2 permission token to the host.
	host <- struct{}{}
	host <- struct{}{}

	// Start dining
	var wg sync.WaitGroup
	for i := 0; i < philosopherCount; i++ {
		wg.Add(1)
		go philosophers[i].eat(&wg)
	}

	// Wait for all philosophers to finish eating
	wg.Wait()
}
