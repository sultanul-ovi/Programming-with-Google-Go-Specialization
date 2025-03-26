package basic

import "fmt"
import "sync"

//lint:ignore U1000 (example)
func concurrent_task(msg string, wg *sync.WaitGroup) {
	defer wg.Done()  // Reduce counter by one
	fmt.Println(msg)
}

//lint:ignore U1000 (example)
func concurrent_waitgroup_example() {
	var wg sync.WaitGroup
	wg.Add(2)  // Increase counter by two
	go concurrent_task("A", &wg)
	go concurrent_task("B", &wg)
	wg.Wait()  // Wait until counter down to zero
	fmt.Println("All task completed")
}

//lint:ignore U1000 (example)
func prod(v1, v2 int, ch chan int) {
	// Channel communication is synchronous (send and receive at same time)
	// Unbuffered channel: sending will block until receive 
	// Buffered channel: sending will block if channel buffer is full
	ch <- v1 * v2
}

//lint:ignore U1000 (example)
func concurrent_channel_example() {
	ch := make(chan int)  // Unbuffered channel (cannot hold data)
	_ = make(chan int, 2)  // Buffered channel (Hold max 2 data)
	go prod(1, 2, ch)
	go prod(3, 4, ch)
	a := <- ch  // Unbuffered channel: blocks if no value is sent
	b := <- ch  // Buffered channel: blocks if buffer is empty
	fmt.Println(a*b)
}

//lint:ignore U1000 (example)
func iterate_channel_example() {
	ch := make(chan string)
	go func() {	
		for message := range ch {
			fmt.Println(message)
		}  // Loop stops when channel is closed
	}()
	ch <- "Hello"
	ch <- "Goodbye"
	close(ch)
}

//lint:ignore U1000 (example)
func select_channel_receive_example() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	defer close(ch1)
	ch2 := make(chan string)
	defer close(ch2)
	go func() {
		// Receive either from channel 1 or 2
		defer wg.Done()
		select {
		case a := <- ch1:
			fmt.Println(a)
		case b := <- ch2:
			fmt.Println(b)
		}
	}()
	wg.Add(1)
	ch2 <- "Hello"
	wg.Wait()
}

//lint:ignore U1000 (example)
func select_channel_send_or_receive_example() {
	var wg sync.WaitGroup
	inchan := make(chan string)
	defer close(inchan)
	outchan := make(chan string)
	defer close(outchan)
	go func() {
		// The one unblocked first will be executed
		// Default: execute if all cases are blocked
		defer wg.Done()
		b := "Goodbye"
		select {
		case a := <- inchan:
			fmt.Println("Received", a)
		case outchan <- b:
			fmt.Println("Sent", b)
		}
	}()
	wg.Add(1)
	<- outchan
	wg.Wait()
}

//lint:ignore U1000 (example)
func select_with_abort_example() {
	var wg sync.WaitGroup
	ch := make(chan string)
	defer close(ch)
	abort := make(chan string)
	defer close(abort)
	go func() {
		defer wg.Done()
		for {
			select {
			case a := <- ch:
				fmt.Println(a)
			case <- abort:
				return
			}
		}
	}()
	wg.Add(1)
	ch <- "Hello"
	ch <- "Goodbye"
	abort <- "Abort"
	wg.Wait()
}

//lint:ignore U1000 (example)
func mutex_example() {
	var wg sync.WaitGroup
	var mut sync.Mutex
	count := 0
	inc := func() {
		defer wg.Done()
		mut.Lock()
		count += 1
		mut.Unlock()
	}
	wg.Add(2)
	go inc()
	go inc()
	wg.Wait()
	fmt.Println(count)
}

//lint:ignore U1000 (example)
func init_example() {
	var wg sync.WaitGroup
	var once sync.Once
	setup := func() {
		fmt.Println("Init")
	}
	dostuff := func() {
		defer wg.Done()
		// Ensure setup() is executed exactly once
		// no matter how many goroutines are called
		once.Do(setup)
		fmt.Println("Hello")
	}
	wg.Add(2)
	go dostuff()  // Init + Hello
	go dostuff()  // Hello
	wg.Wait()
}
