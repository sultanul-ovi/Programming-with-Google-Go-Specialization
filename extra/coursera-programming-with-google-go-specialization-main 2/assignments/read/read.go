package main

import "bufio"
import "fmt"
import "log"
import "os"
import "strings"

const MAX_NAME_LENGTH = 20

type Name struct {
	fname string
	lname string
}

func check_error(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func scan_string() string {
	reader := bufio.NewReader(os.Stdin)
    str, err := reader.ReadString('\n')
	check_error(err)
    str = strings.TrimSpace(str)
	return str
}

func create_name(fname string, lname string) Name {
	if len(fname) > MAX_NAME_LENGTH {
		fname = fname[:MAX_NAME_LENGTH]
	}
	if len(lname) > MAX_NAME_LENGTH {
		lname = lname[:MAX_NAME_LENGTH]
	}
	return Name{fname, lname}
}

func read_names_from_file(input_file string) []Name {
	f, err := os.Open(input_file)
	check_error(err)
	defer f.Close()
	var names []Name
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
    	line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		words := strings.Split(line, " ")
		if len(words) != 2 {
			err = fmt.Errorf("line not containing exactly 2 names: %s", line)
			check_error(err)
		}
		fname, lname := words[0], words[1]
		name := create_name(fname, lname)
		names = append(names, name)
	}
	return names
}

func main() {
	fmt.Print("Enter input file name: ")
	input_file := scan_string()
	names := read_names_from_file(input_file)
	for _, name := range names {
		fmt.Print("First name:", name.fname)
		fmt.Println(", Last name:", name.lname)
	}
}
