package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter several integers(up to 10) split by white space and end with new line:")
	read := bufio.NewReader(os.Stdin)
	line, err := read.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	ints, err := parseInts(line)
	BubbleSort(ints)
	fmt.Print(ints)
}

func parseInts(s string) ([]int, error) {
	var (
		seperateInts = strings.Fields(s)
		ints         = make([]int, len(seperateInts))
		err          error
	)
	for i, integer := range seperateInts {
		ints[i], err = strconv.Atoi(integer)
		if err != nil {
			return nil, err
		}
	}
	return ints, nil
}

func BubbleSort(integers []int) {
	n := len(integers)
	var (
		swapped bool
	)
	for i := 0; i < n-1; i++ {
		swapped = false
		for j := 0; j < n-i-1; j++ {
			if integers[j] > integers[j+1] {
				Swap(integers, j)
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func Swap(integers []int, i int) {
	var temp int
	temp = integers[i]
	integers[i] = integers[i+1]
	integers[i+1] = temp
}
