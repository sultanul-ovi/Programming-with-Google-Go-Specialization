/*
package main

import (

	"fmt"
	"sync"

)

	type Chops struct {
		sync.Mutex
	}

	type Philos struct {
		leftCs, rightCs *Chops
	}

	func hostPermit(i int, p []*Philos, wg *sync.WaitGroup, wg2 *sync.WaitGroup) {
		p[i].leftCs.Lock()
		p[i].rightCs.Lock()
		p[(i+2)%5].leftCs.Lock()
		p[(i+2)%5].rightCs.Lock()
		wg.Done()
		wg2.Done()
	}

	func eat(i, j int, p []*Philos, wg *sync.WaitGroup) {
		fmt.Printf("starting to eat <%d> for %d time(s)", i+1, j)
		p[i].leftCs.Unlock()
		p[i].rightCs.Unlock()
		wg.Done()
	}

	func main() {
		var wg, wg2 sync.WaitGroup
		chops := make([]Chops, 5)
		philos := make([]*Philos, 5)
		for i := 0; i < 5; i++ {
			philos[i] = &Philos{&chops[i], &chops[(i+1)%5]}
		}
		wg.Add(38)

		for i := 1; i < 4; i++ {
			for j := 0; j < 3; j++ {
				wg2.Add(1)
				go hostPermit(j, philos, &wg, &wg2)
				wg2.Wait()
				go eat(j, i, philos, &wg)
				go eat((j+2)%5, i, philos, &wg)
			}
		}

		wg.Wait()
	}
*/
/*package main

import (
	"fmt"
	"sync"
)

type Chops struct {
	sync.Mutex
}

type Philos struct {
	number  int
	leftCs  *Chops
	rightCs *Chops
}

func (p Philos) eat(host chan<- Philos, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		host <- p // Request permission to eat
		fmt.Printf("starting to eat %d\n", p.number)

		p.leftCs.Lock()
		p.rightCs.Lock()

		fmt.Printf("finishing eating %d\n", p.number)
		p.rightCs.Unlock()
		p.leftCs.Unlock()
	}
}

func host(philos []Philos, host chan Philos) {
	for {
		select {
		case philo := <-host:
			fmt.Printf("Permission granted to philosopher %d\n", philo.number)
		}
	}
}

func main() {
	var wg sync.WaitGroup
	chops := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(Chops)
	}

	philos := make([]Philos, 5)
	for i := 0; i < 5; i++ {
		philos[i] = Philos{i + 1, chops[i], chops[(i+1)%5]}
	}

	hostChan := make(chan Philos, 2) // Buffer of 2 to allow 2 philosophers to eat concurrently

	go host(philos, hostChan)

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go philos[i].eat(hostChan, &wg)
	}

	wg.Wait()
}*/

package main

import (
	"fmt"
	"sync"
)

type Chops struct {
	sync.Mutex
}

type Philos struct {
	number          int
	leftCs, rightCs *Chops
	eatCount        int
}

func (p *Philos) eat(wg *sync.WaitGroup, eatPermission <-chan bool, doneEating chan<- *Philos) {
	defer wg.Done()
	for p.eatCount < 3 {
		<-eatPermission // Wait for permission to eat

		p.leftCs.Lock()
		p.rightCs.Lock()

		p.eatCount++
		fmt.Printf("starting to eat %d\n", p.number)

		fmt.Printf("finishing eating %d\n", p.number)
		p.rightCs.Unlock()
		p.leftCs.Unlock()

		doneEating <- p // Notify that eating is finished
	}
}

func host(philos []*Philos, doneEating <-chan *Philos, eatPermission chan<- bool) {
	// Track number of philosophers eating
	eatingPhilosophers := make(map[int]bool)

	for {
		if len(eatingPhilosophers) < 2 {
			// Allow next philosopher to eat if less than 2 are eating
			eatPermission <- true
		}

		// Wait for a philosopher to finish eating
		philo := <-doneEating
		fmt.Printf("Philosopher %d has finished eating\n", philo.number)
		delete(eatingPhilosophers, philo.number)

		// Check if all philosophers are done
		allDone := true
		for _, p := range philos {
			if p.eatCount < 3 {
				allDone = false
				break
			}
		}
		if allDone {
			return // All philosophers have eaten 3 times
		}
	}
}

func main() {
	var wg sync.WaitGroup
	chops := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(Chops)
	}

	philos := make([]*Philos, 5)
	for i := 0; i < 5; i++ {
		philos[i] = &Philos{number: i + 1, leftCs: chops[i], rightCs: chops[(i+1)%5]}
	}

	eatPermission := make(chan bool, 2) // Buffer size 2 to allow up to 2 philosophers to eat concurrently
	doneEating := make(chan *Philos)

	wg.Add(5)
	go host(philos, doneEating, eatPermission)
	for _, philo := range philos {
		go philo.eat(&wg, eatPermission, doneEating)
	}

	wg.Wait()
	fmt.Println("All philosophers have finished eating.")
}
