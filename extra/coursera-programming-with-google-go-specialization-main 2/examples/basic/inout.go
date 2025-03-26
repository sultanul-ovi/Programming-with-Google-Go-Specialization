package basic

import "bufio"
import "fmt"
import "os"
import "strings"

//lint:ignore U1000 (example)
func printing_example() {
	/*	block
		comment	*/
	x := "Minh"
	fmt.Println("Hello ", x)
	fmt.Printf("Hi %s\n", x)
}

//lint:ignore U1000 (example)
func scanning_example() {
	var appleNum, bananaNum int
	fmt.Print("Number of apple & banana: ")  // 4 5
	num, _ := fmt.Scan(&appleNum, &bananaNum)
	fmt.Println("Number of successful scan: ", num)
}

//lint:ignore U1000 (example)
func read_line_as_string() {
	reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    text = strings.TrimSpace(text)  // remove \n
	_ = text
}

//lint:ignore U1000 (example)
func read_all_line_as_string() {
	scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
    	line := scanner.Text()
		_ = line
	}
}

//lint:ignore U1000 (example)
func read_write_file_example() {
	filename := "test.txt"
	data, _ := os.ReadFile(filename)
	_ = os.WriteFile(filename, data, 0777)  // 3rd arg: file perm
}

//lint:ignore U1000 (example)
func file_access_read_example() {
	filename := "test.txt"
	f, err := os.Open(filename)
	if err != nil {
		panic(err)  // do something
	}
	defer f.Close()

	barr := make([]byte, 10)
	num_read_bytes, _ := f.Read(barr)
	_ = num_read_bytes
}

//lint:ignore U1000 (example)
func file_access_write_example() {
	filename := "output.txt"
	f, err := os.Create(filename)
	if err != nil {
		panic(err)  // do something
	}
	defer f.Close()

	barr := []byte{1, 2, 3}
	num_write_bytes, _ := f.Write(barr)
	_ = num_write_bytes
	num_write_bytes, _ = f.WriteString("Hi\n")
	_ = num_write_bytes

	_ = f.Sync()
}
