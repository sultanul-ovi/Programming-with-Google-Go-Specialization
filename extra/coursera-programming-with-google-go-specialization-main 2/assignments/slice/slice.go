package main

import "bufio"
import "fmt"
import "os"
import "slices"
import "strconv"
import "strings"

func read_next_int() (int, bool) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter next integer (or X to quit): ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		if strings.Compare(strings.ToLower(text), "x") == 0 {
			return 0, true
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Println(text, "is not a valid integer")
			continue
		}
		return num, false
	}
}

func slice_insert(sli []int, num int) []int {
	insert_idx, _ := slices.BinarySearch(sli, num)
	sli = append(sli, num)
	for i := len(sli)-1; i > insert_idx; i-- {
		sli[i-1], sli[i] = sli[i], sli[i-1]
	}
	return sli
}

func print_slice(sli []int) {
	fmt.Println("Current sorted list: ", sli)
}

func main() {
	sli := make([]int, 0, 3)
	for {
		num, flag := read_next_int()
		if flag {
			break
		}
		sli = slice_insert(sli, num)
		print_slice(sli)
	}
}
