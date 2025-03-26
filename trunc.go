package main

import "fmt"

func main() {
	fmt.Print("Enter a floating point number: ")
	var x float64
	_, err := fmt.Scanf("%f", &x)
	if err != nil {
		fmt.Println("You did not enter a valid floating point number!")
		return
	}
	var integerNumber = int(x);
	fmt.Printf("Truncated number: %d\n", integerNumber);
}
