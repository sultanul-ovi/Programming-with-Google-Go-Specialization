package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please input a string:")
	input, _ := reader.ReadString('\n') // Read until there is a new line
	input1 := strings.TrimSpace(input)
	input2 := strings.ToLower(input1)
	if strings.HasPrefix(input2, "i") && strings.Contains(input2, "a") && strings.HasSuffix(input2, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}
}
