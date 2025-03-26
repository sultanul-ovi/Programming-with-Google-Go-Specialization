package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	s := make([]int, 3)
	s = s[:0]
	for {
		fmt.Println("Enter an integer, enter 'X' to exit:")
		var input string
		fmt.Scanln(&input)
		if input == string('X') {
			break
		}
		num, _ := strconv.Atoi(input)
		s = append(s, num)
		sort.Ints(s)
		fmt.Println(s)
	}
}
