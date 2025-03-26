package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	//	numbers := []int{5, 1, -7, 3, 10}

	var numbers []int
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Enter a sequence of integers (type 'END' to finish):")

	for {
		fmt.Print("Enter a number: ")
		scanner.Scan()
		input := scanner.Text()

		// Check if the user entered "END"
		if input == "END" {
			break
		}
		
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'END' to finish.")
			continue
		}

		// Add the number to the list
		numbers = append(numbers, number)
	}

	fmt.Println("You entered the following numbers:", numbers)

	BubbleSort(numbers)
	fmt.Println("Sorted ", numbers)

}

func BubbleSort(slide []int) {
	n := len(slide)
	for i := 0; i < n; i++ {
		for j := 1; j < (n - i); j++ {
			if slide[j-1] > slide[j] {
				//swap elements
				Swap(slide, j)
			}
		}
	}
}

func Swap(slide []int, j int) {
	temp := 0
	//swap elements
	temp = slide[j-1]
	slide[j-1] = slide[j]
	slide[j] = temp
}
