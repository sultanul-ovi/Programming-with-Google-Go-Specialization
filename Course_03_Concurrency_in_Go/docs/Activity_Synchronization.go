package main

import (
	"fmt"
	"sync"
)

const (
	numPhilosophers = 5
	numEating       = 2
	numMeals        = 3
)

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	id             int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
	host           *Host
}

type Host struct {
	concurrentEaters chan struct{}
}

func (p *Philosopher) eat() {
	p.leftChopstick.Lock()
	p.rightChopstick.Lock()
	p.host.concurrentEaters <- struct{}{}

	fmt.Printf("Starting to eat %d\n", p.id)
	fmt.Printf("Finishing eating %d\n", p.id)

	<-p.host.concurrentEaters
	p.leftChopstick.Unlock()
	p.rightChopstick.Unlock()
}

func (p *Philosopher) dine(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < numMeals; i++ {
		p.eat()
	}
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = &Chopstick{}
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	var wg sync.WaitGroup

	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{
			id:             i + 1,
			leftChopstick:  chopsticks[i],
			rightChopstick: chopsticks[(i+1)%numPhilosophers],
			host:           &Host{concurrentEaters: make(chan struct{}, numEating)},
		}
		wg.Add(1)
		go philosophers[i].dine(&wg)
	}

	wg.Wait()
}
