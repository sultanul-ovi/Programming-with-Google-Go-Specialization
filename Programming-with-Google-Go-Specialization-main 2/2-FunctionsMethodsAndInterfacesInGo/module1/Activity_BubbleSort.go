package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	numbers := getUserInput()
	if numbers == nil {
		return
	}

	fmt.Println("Unsorted array:", numbers)
	bubbleSort(numbers)
	fmt.Println("Sorted array:", numbers)
}

func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func swap(numbers []int, i int) {
	numbers[i], numbers[i+1] = numbers[i+1], numbers[i]
}

func readInputLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func convertToIntegers(input string) ([]int, error) {
	numStrings := strings.Fields(input)
	numbers := make([]int, 0, len(numStrings))
	for _, numStr := range numStrings {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, fmt.Errorf("invalid input: please enter valid integers")
		}
		numbers = append(numbers, num)
	}

	if len(numbers) > 10 {
		fmt.Println("Input contains more than 10 integers. Please try again.")
		numbers = getUserInput()
	}

	return numbers, nil
}

func getUserInput() []int {
	fmt.Println("Enter up to 10 integers (separated by spaces): ")
	input := readInputLine()

	numbers, err := convertToIntegers(input)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return numbers
}
