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

func splitIntoFour(arr []int) [][]int {
	length := len(arr)
	partLength := length / 4

	subarrays := make([][]int, 4)
	for i := 0; i < 3; i++ {
		subarrays[i] = arr[i*partLength : (i+1)*partLength]
	}
	subarrays[3] = arr[3*partLength:]

	return subarrays
}

func sortSubarray(subarray []int, wg *sync.WaitGroup) {
	sort.Ints(subarray)
	wg.Done()
}

func main() {
	fmt.Println("Please enter your integers separated by space:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	s := strings.Split(strings.TrimSpace(input), " ")
	a := make([]int, 0, len(s))

	for _, str := range s {
		num, _ := strconv.Atoi(str)
		a = append(a, num)
	}

	subarrays := splitIntoFour(a)
	var wg sync.WaitGroup
	wg.Add(4)

	for i := range subarrays {
		go sortSubarray(subarrays[i], &wg)
	}
	wg.Wait()

	// Merge sorted subarrays
	var mergedArray []int
	for i := 0; i < 4; i++ {
		fmt.Printf("Goroutin %d sorted subarray %d: %v", i, i, subarrays[i])
		mergedArray = append(mergedArray, subarrays[i]...)
	}
	sort.Ints(mergedArray) // Final sort

	fmt.Printf("The sorted complete array: %v\n", mergedArray)
}
