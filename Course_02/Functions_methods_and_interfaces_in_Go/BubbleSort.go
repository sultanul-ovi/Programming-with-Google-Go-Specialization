/*Write a Bubble Sort program in Go. The program
should prompt the user to type in a sequence of up to 10 integers. The program
should print the integers out on one line, in sorted order, from least to
greatest. Use your favorite search tool to find a description of how the bubble
sort algorithm works.

As part of this program, you should write a
function called BubbleSort() which
takes a slice of integers as an argument and returns nothing. The BubbleSort()
function should modify the slice so that the elements are in sorted
order.

A recurring operation in the bubble sort algorithm is
the Swap operation which swaps the position of two adjacent elements in the
slice. You should write a Swap() function which performs this operation. Your Swap()
function should take two arguments, a slice of integers and an index value i which
indicates a position in the slice. The Swap() function should return nothing, but it should swap
the contents of the slice in position i with the contents in position i+1.*/

package main

func swap(sli []int, index int) {
	sli[index], sli[index+1] = sli[index+1], sli[index]
}
func BubbleSort(sli []int) {
	n := len(sli)
	for i := 0; i < n-1; i++ {
		for j := 0; j < (n - i - 1); j++ {
			if sli[j] > sli[j+1] {
				swap(sli, j)
			}
		}
	}
}

// func main() {
// 	var input []int
// 	fmt.Print("enter slice of ints")
// 	reader := bufio.NewReader(os.Stdin)
// 	line, _ := reader.ReadString('\n')
// 	tokens := strings.Fields(line)

// 	for _, token := range tokens {
// 		num, err := strconv.Atoi(token)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		input = append(input, num)
// 	}

// 	BubbleSort(input)

// 	fmt.Println("Sorted input:", input)

// }
