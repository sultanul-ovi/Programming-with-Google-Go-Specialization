package main

import (
	"fmt"
)

func main() {
	var floatnum float32
	var intnum int
	fmt.Println("Please give me your floating point number:")
	_, err :=
		fmt.Scan(&floatnum)
	if err != nil {
		fmt.Println("Invalid input, please make sure it is a floating point number.")
	} else {
		intnum = int(floatnum)
		fmt.Println(intnum)
	}
}
