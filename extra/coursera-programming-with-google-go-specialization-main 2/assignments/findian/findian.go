package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputvalue := scanner.Text()
	value := strings.ToLower(inputvalue)
	if (strings.HasPrefix(value, "i") && strings.ContainsAny(value, "a") && strings.HasSuffix(value, "n")) {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
