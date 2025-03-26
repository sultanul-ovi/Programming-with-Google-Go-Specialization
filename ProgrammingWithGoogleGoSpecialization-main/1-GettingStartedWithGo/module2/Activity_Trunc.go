package main

import "fmt"

func trunc(f float64) {
	fmt.Printf("user input: %d\n", int(f))
}

func main() {
	var f float64
	fmt.Print("Enter a float value: ")
	_, err := fmt.Scanf("%f", &f)
	if err != nil {
		fmt.Println(err)
	}
	trunc(f)
}
