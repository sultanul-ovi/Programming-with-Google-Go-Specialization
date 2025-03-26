package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name string
	var address string
	myMap := make(map[string]string)

	fmt.Print("Enter a name")
	fmt.Scan(&name)
	fmt.Print("Enter an address")
	fmt.Scan(&address)
	myMap["name"] = name
	myMap["address"] = address
	j, err := json.Marshal(myMap)
	if err == nil {
		fmt.Print(string(j))
		return
	}
	fmt.Print(err)

}
