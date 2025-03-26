package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter name: ")
	name := readInputLine()

	fmt.Print("Enter address: ")
	address := readInputLine()

	idMap := make(map[string]string)
	idMap[name] = address

	jsonObj, err := json.Marshal(idMap)
	if err != nil {
		fmt.Println("Error converting to JSON:", err)
		return
	}

	fmt.Println("JSON Object:")
	fmt.Println(string(jsonObj))
}

func readInputLine() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}
