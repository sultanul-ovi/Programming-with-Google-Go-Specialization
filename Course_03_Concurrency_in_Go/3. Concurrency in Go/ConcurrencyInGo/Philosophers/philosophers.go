package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

type Chopstick struct {
	id  int
	mut sync.Mutex
}
type Philosopher struct {
	id    int
	eaten int
	Left  *Chopstick
	Right *Chopstick
}

type Host struct {
	CurrentEating int
}

func (p Philosopher) Eating(host *Host, wg *sync.WaitGroup) {
	defer wg.Done()

	for p.eaten < 3 {
		if host.CurrentEating < 2 {
			host.CurrentEating++

			trial := rand.Intn(2)
			if trial == 0 {
				p.Left.mut.Lock()
				p.Right.mut.Lock()
			} else {
				p.Right.mut.Lock()
				p.Left.mut.Lock()
			}
			fmt.Println("Starting to eat " + strconv.Itoa(p.id))
			p.Left.mut.Unlock()
			p.Right.mut.Unlock()
			p.eaten++
			fmt.Println("Finishing eating " + strconv.Itoa(p.id))
		}
		host.CurrentEating--

	}
}
func main() {
	host := Host{0}
	philosophers := make([]Philosopher, 5)

	for i := 1; i < 6; i++ {
		philosophers[i-1] = Philosopher{id: i}
	}

	chopsticks := make([]Chopstick, 5)
	for i := 0; i < 5; i++ {
		chopsticks[i] = Chopstick{id: i}
	}

	for i := range len(chopsticks) {
		philosophers[i].Left = &chopsticks[i]
		philosophers[i].Right = &chopsticks[(i+1)%5]
	}
	var wg sync.WaitGroup
	wg.Add(5)

	for i := 0; i < 5; i++ {
		go philosophers[i].Eating(&host, &wg)
	}
	wg.Wait()
}
