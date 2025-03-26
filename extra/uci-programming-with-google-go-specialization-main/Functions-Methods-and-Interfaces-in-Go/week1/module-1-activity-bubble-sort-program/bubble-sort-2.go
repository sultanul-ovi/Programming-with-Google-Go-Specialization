package main

import (
	"fmt"
	"strconv"
)

func swap(i int, sli []int) {
	var curr, next int
	curr = sli[i]
	next = sli[i+1]
	if curr > next {
		sli[i] = next
		sli[i+1] = curr
	}
}

func bubble_sort(slice_to_sort []int) {
	fmt.Printf("Sorting ==> %v\n", slice_to_sort)
	rounds := 1
	for {
		swap_counter := 0
		for i := 0; i < len(slice_to_sort)-1; i++ {
			if slice_to_sort[i] > slice_to_sort[i+1] {
				swap(i, slice_to_sort)
				swap_counter++
			}
		}
		rounds++
		if swap_counter == 0 {
			break
		}
	}
	fmt.Printf("Sorted array ==> %v -- (Number of rounds needed: %d)\n", slice_to_sort, rounds)
}

func read_in_slice_to_sort() []int {
	var s string
	input_slice := make([]int, 0, 10)
	for {
		fmt.Printf("Enter %v more integers -- or press x! ", 10-len(input_slice))
		fmt.Scanf("%s", &s)
		if s == "x" || s == "X" {
			return input_slice
		}
		number, err := strconv.Atoi(s)
		if err != nil {
			continue
		} else {
			input_slice = append(input_slice, number)
		}
		if len(input_slice) == 10 {
			return input_slice
		}
	}
	return input_slice
}

func main() {
	slice_to_sort := read_in_slice_to_sort()
	bubble_sort(slice_to_sort)
}
