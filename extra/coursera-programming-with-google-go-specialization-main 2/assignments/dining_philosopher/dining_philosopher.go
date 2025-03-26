package main

import (
	"fmt"
	"sync"
)

const NumPeople = 5
const NumMeals = 3
const QueueSize = 2

type Null struct{}

var null Null = struct{}{}

type Chopstick struct {
	sync.Mutex
}

type Philosopher struct {
	number int
	leftCS, rightCS *Chopstick
}

func (p Philosopher) eat(permCh, queueCh chan Null, wg *sync.WaitGroup) {
	defer wg.Done()
	for range NumMeals {
		p.leftCS.Lock()
		p.rightCS.Lock()
		<- permCh  // Ask permission to eat
		fmt.Println("starting to eat", p.number)
		fmt.Println("finishing eating", p.number)
		<- queueCh // Leave the queue
		p.leftCS.Unlock()
		p.rightCS.Unlock()
	}
}

type Host struct{}

func (h Host) host(permCh, queueCh chan Null, wg *sync.WaitGroup) {
	defer wg.Done()
	for range (NumPeople * NumMeals) {
		queueCh <- null // Wait until queue is not full
		permCh <- null  // Grant permission to eat
	}
}

func main() {
	chopsticks := make([]*Chopstick, NumPeople)
	for i := range NumPeople {
		chopsticks[i] = &Chopstick{}
	}
	philosophers := make([]*Philosopher, NumPeople)
	for i := range NumPeople {
		philosophers[i] = &Philosopher{i+1, chopsticks[i], chopsticks[(i+1)%NumPeople]}
	}
	host := &Host{}
	permCh := make(chan Null)
	queueCh := make(chan Null, QueueSize)
	var wg sync.WaitGroup
	wg.Add(NumPeople+1)
	go host.host(permCh, queueCh, &wg)
	for i := range NumPeople {
		go philosophers[i].eat(permCh, queueCh, &wg)
	}
	wg.Wait()
}
