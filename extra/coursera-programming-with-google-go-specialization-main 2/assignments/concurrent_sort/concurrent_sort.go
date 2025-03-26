package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"sync"
)

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
    line := scan_line()
	var int_slice []int
	for _, field := range strings.Fields(line) {
		num, err := strconv.Atoi(field)
		if err != nil {
			log.Fatal(err)
		}
		int_slice = append(int_slice, num)
	}
	return int_slice
}

func sort_portion(int_portion []int, wg *sync.WaitGroup) {
	defer wg.Done()
	old_portion := make([]int, len(int_portion))
	copy(old_portion, int_portion)
	slices.Sort(int_portion)
	fmt.Println("Portion sorted:", old_portion, "=>", int_portion)
}

func merge_sorted_portions(int_portions []int, split_idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := len(int_portions)
	n1, n2 := split_idx, n-split_idx
	old_portion_1 := make([]int, n1)
	old_portion_2 := make([]int, n2)
	copy(old_portion_1, int_portions[:split_idx])
	copy(old_portion_2, int_portions[split_idx:])
	i, i1, i2 := 0, 0, 0
	for i < n {
		if i1 >= n1 {
			int_portions[i] = old_portion_2[i2]
			i += 1
			i2 += 1
			continue
		}
		if i2 >= n2 {
			int_portions[i] = old_portion_1[i1]
			i += 1
			i1 += 1
			continue
		}
		if old_portion_1[i1] <= old_portion_2[i2] {
			int_portions[i] = old_portion_1[i1]
			i += 1
			i1 += 1
		} else {
			int_portions[i] = old_portion_2[i2]
			i += 1
			i2 += 1
		}
	}
	fmt.Println("Merge sorted:", old_portion_1, old_portion_2, "=>", int_portions)
}

func main() {
	fmt.Print("Enter space-separated integers to be sorted: ")
	int_slice := scan_int_slice()
	n := len(int_slice)
	n1, n2, n3 := n/4, 2*n/4, 3*n/4
	var wg sync.WaitGroup
	wg.Add(4)
	go sort_portion(int_slice[0:n1], &wg)
	go sort_portion(int_slice[n1:n2], &wg)
	go sort_portion(int_slice[n2:n3], &wg)
	go sort_portion(int_slice[n3:n], &wg)
	wg.Wait()
	wg.Add(2)
	go merge_sorted_portions(int_slice[0:n2], n1, &wg)     // Merge the first two portions
	go merge_sorted_portions(int_slice[n2:n], n3-n2, &wg)  // Merge the last two portions
	wg.Wait()
	wg.Add(1)
	go merge_sorted_portions(int_slice, n2, &wg)  // Merge the final two halfs
	wg.Wait()
	fmt.Println("Final sorted list:", int_slice)
}
