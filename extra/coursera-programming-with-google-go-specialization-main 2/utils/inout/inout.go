package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

//lint:ignore U1000 (example)
func scan_line() string {
	scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	line := scanner.Text()
	return line
}

//lint:ignore U1000 (example)
func scan_all_lines() []string {
	scanner := bufio.NewScanner(os.Stdin)
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
	return lines
}

//lint:ignore U1000 (example)
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

func main() {
	// Scan line as string
	fmt.Print("Input string (can contain whitespace): ")
	str := scan_line()
	fmt.Println("Your string is:", str)
}
