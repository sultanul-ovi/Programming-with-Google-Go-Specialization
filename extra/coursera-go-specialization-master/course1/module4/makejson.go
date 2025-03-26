package main

import (
"encoding/json"
"fmt" )
	

func main() {
		
	var name string
	fmt.Print("Enter name: ")
	fmt.Scan(&name)

	var address string
	fmt.Print("Enter address: ")
	fmt.Scan(&address)

	var person map[string]string = make(map[string]string)
	person["name"] = name
	person["address"] = address
	
	json, _ := json.Marshal(person)

	fmt.Println(string(json))	
	

}
