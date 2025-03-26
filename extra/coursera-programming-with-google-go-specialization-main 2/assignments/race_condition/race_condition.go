// Race condition explaination:
// The user has a budget of 100$. Two transaction (50$ and 70$) are executed concurrently.
// Each transaction check if there is enough budget before proceeding the transaction that
// takes 1 second to execute. As the two goroutine has read and write access to same variable,
// it led to the case where the user buy both games even when there is not enough budget.
package main

import "fmt"
import "sync"
import "time"

var budget int = 100
var wg sync.WaitGroup

func buy_zelda() {
	defer wg.Done()
	price := 50
	if budget < price {
		fmt.Println("Zelda transaction declined")
		return
	}
	fmt.Println("Waiting for Zelda transaction")
	time.Sleep(1 * time.Second)
	budget -= price
	fmt.Println("Zelda transaction completed")
}

func buy_pokemon() {
	defer wg.Done()
	price := 70
	if budget < price {
		fmt.Println("Pokemon transaction declined")
		return
	}
	fmt.Println("Waiting for Pokemon transaction")
	time.Sleep(1 * time.Second)
	budget -= price
	fmt.Println("Pokemon transaction completed")
}

func main() {
	wg.Add(2)
	go buy_zelda()
	go buy_pokemon()
	wg.Wait()
	fmt.Println("Remaining budget:", budget)
}
