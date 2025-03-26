package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	// Initialize an empty slice with a length of 3
	numbers := make([]int, 0, 3)

	fmt.Println("Enter integers to add to the slice. Type 'X' to quit.")

	var input string

	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")

		fmt.Scan(&input)

		// Check if the user entered 'X' to quit
		if input == "X" {
			fmt.Println("Quiting the program.")
			break
		}

		// Convert input to an integer
		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Enter an integer or 'X' to quit.")
			continue
		}

		// Add the number to the slice
		numbers = append(numbers, num)

		// Sort the slice
		sort.Ints(numbers)

		// Print the sorted slice
		fmt.Println("Sorted slice:", numbers)
	}
}
