package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	sli := make([]int, 0, 4)
	for {
		var input string
		fmt.Println("Please enter a number, or enter X if you want to exit:")
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Invalid input.")
		}
		if input == "X" {
			break
		} else {
			var input1 int
			input1, err = strconv.Atoi(input)
			if err != nil {
				fmt.Println("Invalid input.")
			} else {
				sli = append(sli, input1)
				sort.Ints(sli)
				fmt.Println(sli)
			}
		}
	}
	fmt.Println(sli)
}
