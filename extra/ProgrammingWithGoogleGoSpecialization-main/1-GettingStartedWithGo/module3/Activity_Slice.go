package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)

	for {
		fmt.Print("Enter an integer or X to exit the program: ")
		var userinput string
		fmt.Scan(&userinput)
		userinput = strings.ToLower(userinput)
		if userinput == "x" {
			break
		}
		num, err := strconv.Atoi(userinput)
		if err != nil {
			fmt.Printf("Invalid input: %s, err: %s\n", userinput, err)
			continue
		}
		slice = append(slice, num)
		sort.Ints(slice)

		fmt.Println("Sorted slice: ", slice)
	}

}
