package main

import "bufio"
import "encoding/json"
import "fmt"
import "log"
import "os"
import "strings"

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

func main() {
	fmt.Print("Input name: ")
	name := scan_string()
	fmt.Print("Input address: ")
	address := scan_string()
	data := map[string]string{"name": name, "address": address}
	json_obj, err := json.Marshal(data)
	check_error(err)
	fmt.Println(string(json_obj))
}
