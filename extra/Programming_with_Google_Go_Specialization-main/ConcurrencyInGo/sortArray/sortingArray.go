package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func merge(arr1, arr2 []int) []int {
	i, j := 0, 0
	result := make([]int, 0, len(arr1)+len(arr2))

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result = append(result, arr1[i])
			i++
		} else {
			result = append(result, arr2[j])
			j++
		}
	}
	result = append(result, arr1[i:]...)
	result = append(result, arr2[j:]...)
	return result
}
func main() {
	//create a scanner to read input from stdin
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter comma seperated list of numbers. Total should be divisible by 4")
	//Read user input
	scanner.Scan()
	input := scanner.Text()

	//split input by commas
	parts := strings.Split(input, ",")
	var nums []int

	for _, num := range parts {
		num, err := strconv.Atoi(strings.TrimSpace(num))
		if err != nil {
			fmt.Println("invalid input:", num)
			return
		}
		nums = append(nums, num)
	}
	fmt.Println("parsed input integers and inserted into slice nums")

	n := len(nums)
	if n < 4 {
		fmt.Println("Not enough numbers")
		return
	}
	//determine partition size
	partitionSize := (n + 3) / 4

	//create 4 partitions
	partitions := make([][]int, 4)
	for i := 0; i < 4; i++ {
		start := i * partitionSize
		end := start + partitionSize
		if end > n {
			end = n
		}
		partitions[i] = nums[start:end]
	}

	//use waitgroup to wait for all goroutines

	var wg sync.WaitGroup
	wg.Add(4)

	for i := 0; i < 4; i++ {
		go func(index int) {
			defer wg.Done()
			fmt.Println("Goroutine sorting partition:", partitions[index])
			sort.Ints(partitions[index])
			fmt.Println("Sorted partition:", partitions[index])
		}(i)
	}
	//wait for all sorting goroutines to complete
	wg.Wait()

	//merge sorted partitions sequentially
	sortedArray := merge(merge(partitions[0], partitions[1]), merge(partitions[2], partitions[3]))
	fmt.Println("Sorted array:", sortedArray)
}
