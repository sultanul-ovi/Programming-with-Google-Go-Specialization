package main

import (
	"fmt"
	"strings"
)

func findian(str string) bool {
	lower := strings.ToLower(str)
	return strings.HasPrefix(lower, "i") && strings.HasSuffix(lower, "n") && strings.Contains(str, "a")
}

func main() {
	var str string
	fmt.Print("Enter a string: ")
	_, err := fmt.Scanf("%s", &str)
	if err != nil {
		fmt.Println(err)
	}
	if findian(str) {
		fmt.Println("FOUND")
	} else {
		fmt.Println("NOT FOUND")
	}
}
