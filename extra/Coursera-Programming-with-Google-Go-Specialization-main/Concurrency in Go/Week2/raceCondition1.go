// I create two function and use "go" to assign them to different threads.
// And due to the race condition, we can not predict whether the "add" function be excuted first or the "division" function, and also we can know when the function read which parameter.
// If you run enough times, you can find that the result change.

package main

import (
	"fmt"
	"time"
)

func add(i *int) {
	(*i)++
	fmt.Println(*i)
}

func division(i *int) {
	(*i) = (*i) / 2
	fmt.Println(*i)
}

func main() {
	var i int = 2
	go add(&i)
	go add(&i)
	go division(&i)
	go add(&i)
	time.Sleep(time.Second) //Add this to prevent the function quit without printing the result. Since the main function doesn't wait for the goroutine to finish its execution. In Go, when the main function returns, the program exits, and all goroutines are killed.
}
