package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Swap a[i] and a[i+1] in place
func Swap(a []int, i int) {
	a[i], a[i+1] = a[i+1], a[i]
}

// Bubble sort integers slice in place
func BubbleSort(a []int) {
	n := len(a)
	for i := n-1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if a[j] > a[j+1] {
				Swap(a, j)
			}
		}
	}
}

func scan_line() string {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	line := scanner.Text()
	return line
}

func scan_int_slice() []int {
	num_limit := 10
	var int_slice []int
	received_input := false
	for !received_input {
		fmt.Print("Enter up to 10 integers (space separated): ")
		line := scan_line()
		int_slice = make([]int, 0)
		received_input = true
		for idx, field := range strings.Fields(line) {
			num, err := strconv.Atoi(field)
			if err != nil {
				fmt.Println("Input contains invalid integers")
				received_input = false
				break
			}
			if idx >= num_limit {
				fmt.Println("Input contains more than 10 integers")
				received_input = false
				break
			}
			int_slice = append(int_slice, num)
		}
	}
	return int_slice
}

func print_int_slice(a []int) {
	fmt.Println("Sorted slice:", a)
}

func main() {
	a := scan_int_slice()
	BubbleSort(a)
	print_int_slice(a)
}
