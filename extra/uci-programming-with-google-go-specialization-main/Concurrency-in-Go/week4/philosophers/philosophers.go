package main

import (
	"fmt"
	"sync"
)

/**
Implement the dining philosopher’s problem with the following constraints/modifications.

1.There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
2.Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)
3.The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
4.In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
5.The host allows no more than 2 philosophers to eat concurrently.
6.Each philosopher is numbered, 1 through 5.
7.When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.
8.When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.
*/

const numPhilosophers = 5
const numMeals = 3 // 每位哲学家只吃3次饭

var (
	hostChannel = make(chan struct{}, 2)
	wg          sync.WaitGroup
)

type Chopstick struct{ sync.Mutex }

type Philosopher struct {
	number         int
	leftChopstick  *Chopstick
	rightChopstick *Chopstick
}

func (p Philosopher) eat() {
	for i := 0; i < numMeals; i++ {
		hostChannel <- struct{}{}
		p.leftChopstick.Lock()
		p.rightChopstick.Lock()

		fmt.Printf("starting to eat %v\n", p.number)
		fmt.Printf("finishing eating %v\n", p.number)

		p.rightChopstick.Unlock()
		p.leftChopstick.Unlock()

		<-hostChannel
	}
	wg.Done()
}

func main() {
	chopsticks := make([]*Chopstick, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		chopsticks[i] = new(Chopstick)
	}

	philosophers := make([]*Philosopher, numPhilosophers)
	for i := 0; i < numPhilosophers; i++ {
		philosophers[i] = &Philosopher{i + 1, chopsticks[i], chopsticks[(i+1)%numPhilosophers]}
	}

	for i := 0; i < numPhilosophers; i++ {
		wg.Add(1)
		go philosophers[i].eat()
	}

	wg.Wait()
}
