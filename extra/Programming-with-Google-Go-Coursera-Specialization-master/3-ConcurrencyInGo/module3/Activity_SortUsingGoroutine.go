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

type Result struct {
	index     int
	partition []int
}

func main() {
	// user prompt
	input := userPrompt("Enter a series of integers (separated by spaces):")

	// Split the input into an array of integers
	integers, err := parseIntArray(input)
	if err != nil {
		fmt.Println("Error parsing integers:", err)
		return
	}

	// split parsed user input into 2d array (arrayLength by numPartitions)
	numPartitions := 4
	partitions := divideIntegers(integers, numPartitions)

	// use channels and wait group to sync sorting goroutines
	var wg sync.WaitGroup
	resultCh := make(chan Result, numPartitions)

	// retrieve and merge sorted partitions
	sortedIntegers := sortAndMergePartitions(partitions, &wg, resultCh)

	fmt.Println("sorted list:")
	fmt.Println(sortedIntegers)
}

func userPrompt(prompt string) string {
	fmt.Println(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	return input
}

// parseIntArray parses a space-separated string of integers into an array of integers.
func parseIntArray(input string) ([]int, error) {
	parts := strings.Fields(input)
	integers := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		integers[i] = num
	}

	return integers, nil
}

// divideIntegers divides an array of integers into numPartitions subarrays
func divideIntegers(integers []int, numPartitions int) [][]int {
	partitionSize := len(integers) / numPartitions
	partitions := make([][]int, numPartitions)
	for i := 0; i < numPartitions; i++ {
		start := i * partitionSize
		end := (i + 1) * partitionSize
		if i == numPartitions-1 {
			end = len(integers)
		}
		partitions[i] = integers[start:end]
	}
	return partitions
}

// sortIntegers sorts an array of integers
func sortIntegers(integers []int) []int {
	sort.Ints(integers)
	return integers
}

// sortAndMergePartitions sorts and merges partitions concurrently
func sortAndMergePartitions(partitions [][]int, wg *sync.WaitGroup, resultCh chan Result) []int {
	var sortedIntegers []int

	// sort each partition concurrently
	for i, partition := range partitions {
		wg.Add(1)
		go func(partition []int, idx int) {
			defer wg.Done()
			sortedPartition := sortIntegers(partition)
			fmt.Printf("sorted partition %d: %v\n", idx+1, sortedPartition)
			resultCh <- struct {
				index     int
				partition []int
			}{idx, sortedPartition}
		}(partition, i)
	}

	// close the result channel when all sorting goroutines are done
	go func() {
		wg.Wait()       // Wait for all sorting goroutines to finish
		close(resultCh) // close the result channel
	}()

	// collect the sorted partitions and merge them
	sortedPartitions := make([][]int, len(partitions))
	for i := 0; i < len(partitions); i++ {
		result := <-resultCh
		sortedPartitions[result.index] = result.partition
	}

	for _, sortedPartition := range sortedPartitions {
		sortedIntegers = merge(sortedIntegers, sortedPartition)
	}

	return sortedIntegers
}

// merge merges two separate arrays into one
func merge(left []int, right []int) []int {
	result := make([]int, 0)
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}
