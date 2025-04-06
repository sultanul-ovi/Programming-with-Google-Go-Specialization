package classic

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
	eatCount int
}

func (p Philosopher) eat(wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eatCount < eatTimes {
		// Request permission from the host
		//p.host <- struct{}{}

		// Pick up chopsticks in any order

		if p.number != philosopherCount {
			p.left.Lock()
			fmt.Printf("Philosopher %d tooked the left \n", p.number)
			p.right.Lock()
			fmt.Printf("Philosopher %d tooked the right \n", p.number)
		} else {
			// El ultimo filosofo empieza cogiendo el de su derecha para evitar la dependencia circular
			p.right.Lock()
			fmt.Printf("Philosopher %d tooked the right \n", p.number)
			p.left.Lock()
			fmt.Printf("Philosopher %d tooked the left \n", p.number)
		}

		// Start eating
		fmt.Printf("starting to eat %d\n", p.number)
		time.Sleep(time.Millisecond * 100) // Simulate eating time

		// Finish eating
		fmt.Printf("finishing eating %d\n", p.number)
		p.eatCount++

		// Put down chopsticks
		p.right.Unlock()
		p.left.Unlock()

		// Release the permission
		//<-p.host
	}
}

func doIt() {

	chopsticks := make([]sync.Mutex, philosopherCount)
	philosophers := make([]*Philosopher, philosopherCount)

	for i := 0; i < philosopherCount; i++ {
		philosophers[i] = &Philosopher{
			number:   i + 1,
			right:    &chopsticks[i],
			left:     &chopsticks[(i+1)%philosopherCount],
			eatCount: 0}
	}

	var wg sync.WaitGroup
	for i := 0; i < philosopherCount; i++ {
		wg.Add(1)
		go philosophers[i].eat(&wg)
	}

	// Wait for all philosophers to finish eating
	wg.Wait()	
}
