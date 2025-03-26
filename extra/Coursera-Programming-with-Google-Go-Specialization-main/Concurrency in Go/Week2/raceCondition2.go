// The same as race condition 1, we divided the code into four threads, and we can not determine which one to be excuted first or later, so the results is unpredictable.
// If you run this program for several times, you will find the results change.
// If the print func be excuated first, the result will be "2 3", otherwise it will be "5 3"
package main

import (
	"fmt"
	"time" // Use this package to control the begin and end of goroutines
)

var x, y int = 2, 3

func main() {
	go func() {
		x = x + y
	}()
	go func() {
		fmt.Println(x, y)
	}()
	time.Sleep(time.Second)
}
