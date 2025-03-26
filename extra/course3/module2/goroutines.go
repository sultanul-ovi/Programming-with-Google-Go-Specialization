package main

import (
	"fmt"
	"time"
)


var activeConnections = 0

func main() {
	
	// This is a example that mimic a web server. Imagine that for every HTTP request the method newHttpRequest
	// gets called and increases a counter which represents active connections. At the end of the method
	// when the method is done decreases the counter.
	// When all the request are done shouldn't be any activeConnections (activeConnections = 0 ). Unfortunatly doesn't go to 0 because of 
	// a race condition.

	numberOfRequest := 2000
	
	for i:= 0 ; i < numberOfRequest; i ++{
		go newHttpRequest()
	}	
	time.Sleep(3 * time.Second)	
	fmt.Println("Number of active connections is " , activeConnections)
}

func newHttpRequest() {	
	activeConnections = activeConnections + 1		
	time.Sleep(100 * time.Millisecond)	// Do some heavy work
	activeConnections = activeConnections - 1 
}
