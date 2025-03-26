package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	type Name struct {
		fname string
		lname string
	}

	sli := make([]Name, 0, 100)

	var filename string
	fmt.Println("Please enter your file name with suffix '.txt', and make sure your file is in the same path of this file:")
	fmt.Scan(&filename)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Your file name should end with '.txt'. and you should put it in the same folder as this file")
	} else {
		fileScanner := bufio.NewScanner(file)
		fileScanner.Split(bufio.ScanLines)
		for fileScanner.Scan() {
			var line string
			var fullname Name
			line = fileScanner.Text()
			fullnamearr := strings.Fields(line)
			fullname.fname = fullnamearr[0]
			fullname.lname = fullnamearr[1]
			sli = append(sli, fullname)
		}
	}
	fmt.Println(sli)
}
