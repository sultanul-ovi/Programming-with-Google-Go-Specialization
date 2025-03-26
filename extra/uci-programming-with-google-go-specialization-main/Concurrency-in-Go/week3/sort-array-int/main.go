package main

import (
	"fmt"
	"sort"
)

/**
编写一个对整数数组排序的程序。
程序应将数组划分为 4 个部分，每个部分由不同的程序进行排序。
每个分区的大小应大致相等。
然后，主程序应将 4 个已排序的子数组合并为一个大的已排序数组。

程序应提示用户输入一系列整数。
每个对数组的 1/4 进行排序的 goroutine 都应打印它要排序的子数组。
排序完成后，主程序应打印整个排序列表。
*/

func sortPartition(arr []int, ch chan []int) {
	sort.Ints(arr)
	ch <- arr
}

func mergeArrays(arr1, arr2 []int) []int {
	result := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			result[k] = arr1[i]
			i++
		} else {
			result[k] = arr2[j]
			j++
		}
		k++
	}
	for i < len(arr1) {
		result[k] = arr1[i]
		i++
		k++
	}
	for j < len(arr2) {
		result[k] = arr2[j]
		j++
		k++
	}
	return result
}

func main() {
	var arr []int
	num := 0

	fmt.Println("Enter integers (enter -1 to finish):")
	for {
		fmt.Scanln(&num)
		if num == -1 {
			break
		}
		arr = append(arr, num)
	}

	ch := make(chan []int, 4)
	partSize := len(arr) / 4

	for i := 0; i < 4; i++ {
		start := i * partSize
		end := (i + 1) * partSize
		if i == 3 {
			end = len(arr)
		}
		go sortPartition(arr[start:end], ch)
		fmt.Printf("Partition %d to be sorted: %v\n", i+1, arr[start:end])
	}

	var results [][]int
	for i := 0; i < 4; i++ {
		results = append(results, <-ch)
	}

	sortedArr := mergeArrays(mergeArrays(results[0], results[1]), mergeArrays(results[2], results[3]))

	fmt.Println("Sorted Array:")
	fmt.Println(sortedArr)
}
