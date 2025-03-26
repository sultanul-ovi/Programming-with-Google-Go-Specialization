package main

import "fmt"

func main() {
	fmt.Println("Enter a float point number:")

    var decimalNumber float32
    fmt.Scan(&decimalNumber)
           
    fmt.Printf("%d", int(decimalNumber))

}
