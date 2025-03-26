package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	fname [20]byte
	lname [20]byte
}

func main() {
	var txtFile string
	fmt.Print("Enter txt file to open: ")
	_, err := fmt.Scanf("%s", &txtFile)
	if err != nil {
		fmt.Printf("invalid user input: %v", err)
	}

	f, err := os.Open(txtFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	var names []Name

	s := bufio.NewScanner(f)

	for s.Scan() {
		line := s.Text()
		nameFields := strings.Fields(line)
		if len(nameFields) >= 2 {
			var fname [20]byte
			var lname [20]byte

			copy(fname[:], nameFields[0])
			copy(lname[:], nameFields[1])

			names = append(names, Name{fname, lname})
		}
	}

	if err := s.Err(); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("First and Last Names:")
	for _, name := range names {
		fmt.Printf("%s %s\n", name.fname, name.lname)
	}
}
