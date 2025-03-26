package main

import "fmt"

/**
用 Go 编写一个冒泡排序程序。
程序应提示用户输入最多 10 个整数的序列。
程序应按从最小到最大的排序顺序将整数打印在一行上。

作为该程序的一部分，您应编写一个名为名为 BubbleSort() 的函数，
该函数以整数片段为参数，不返回任何内容。BubbleSort() 函数应修改片段，使元素按排序顺序排列。

冒泡排序算法中经常出现的操作是交换操作，它可以交换相邻两个元素在切片中的位置。
切片中相邻两个元素的位置。您应该编写一个 Swap() 函数来执行该操作。
您的 Swap()函数应接受两个参数，即一个整数片段和一个索引值 i（表示在片段中的位置）。
Swap() 函数不会返回任何内容，但会交换i+1 位置的内容。
*/

func main() {
	var n int
	nums := make([]int, 0, 10)

	fmt.Print("Please enter the number of integers (up to 10):")
	fmt.Scan(&n)

	if n > 10 {
		fmt.Println("Error: you can enter a maximum of 10 integers.")
		return
	}

	fmt.Println("Please enter integers:")
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		nums = append(nums, num)
	}

	// 执行冒泡排序
	BubbleSort(nums)

	fmt.Println(nums)
}

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 如果左边的数大于右边的数，交换
				Swap(arr, j)
			}
		}
	}
}

// Swap 函数交换切片中索引 i 和 i+1 处的元素。
func Swap(arr []int, i int) {
	arr[i], arr[i+1] = arr[i+1], arr[i]
}
