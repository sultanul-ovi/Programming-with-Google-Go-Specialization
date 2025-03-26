package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func mergeTwoSortedArrays(a, b []int) []int {
	merged := []int{}

	var i int = 0
	var j int = 0

	for i < len(a) && j < len(b) {

		// Find the smaller of the two current elements and move ahead in the array of the picked element
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	// if there are remaining elements of the first array, add them
	for i < len(a) {
		merged = append(merged, a[i])
		i++
	}

	// if there are remaining elements of the second array, add them
	for j < len(b) {
		merged = append(merged, b[j])
		j++
	}
	return merged

}

// mergeSortedArrays merges multiple sorted slices into one sorted slice.
func mergeSortedArrays(parts [][]int) []int {
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result = mergeTwoSortedArrays(result, parts[i])
	}
	return result
}

func sortSlice(slice []int, wg *sync.WaitGroup, subroutineNumber int) {
	defer wg.Done()

	fmt.Printf("Subroutine %d holds: %v", subroutineNumber, slice)
	fmt.Println()
	sort.Ints(slice)
	fmt.Printf("Subroutine %d sorted: %v", subroutineNumber, slice)
	fmt.Println()
}

func main() {

	var arr []int
	var input string

	fmt.Println("Enter a sequence of integers (type 'END' to finish):")

	for {
		fmt.Print(">")
		fmt.Scan(&input)

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
		arr = append(arr, number)

	}

	fmt.Println("Unsorted array:", arr)

	n := len(arr)
	if n < 4 {
		fmt.Println("Array is too small to partition into 4 parts. Sorting directly.")
		sort.Ints(arr)
		fmt.Println("Sorted Array:", arr)
		return
	}

	// Determine partition sizes
	partSize := n / 4
	rem := n % 4

	// Create partitions
	partitions := make([][]int, 4)
	start := 0
	for i := 0; i < 4; i++ {
		end := start + partSize
		if i < rem {
			end++
		}
		partitions[i] = arr[start:end]
		start = end
	}

	// Sort each partition in parallel using goroutines
	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go sortSlice(partitions[i], &wg, i+1)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	fmt.Println("Sorted Partitions:", partitions)

	// Merge all sorted partitions into one sorted array
	sortedArray := mergeSortedArrays(partitions)
	fmt.Println("Sorted Array:", sortedArray)

}
