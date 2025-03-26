package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var idMap map[string]string
	idMap = make(map[string]string)
	var name string
	var address string
	fmt.Println("Please enter the name:")
	_, err1 := fmt.Scan(&name)
	if err1 != nil {
		fmt.Println("Invalid value.")
	} else {
		idMap["name"] = name
		fmt.Println("Please enter the address:")
		_, err2 := fmt.Scan(&address)
		if err2 != nil {
			fmt.Println("Invalid value.")
		} else {
			idMap["address"] = address
			idJson, err3 := json.Marshal(idMap)
			if err3 != nil {
				fmt.Println("Convertion failed.")
			} else {
				fmt.Println(string(idJson))
			}
		}
	}

}
