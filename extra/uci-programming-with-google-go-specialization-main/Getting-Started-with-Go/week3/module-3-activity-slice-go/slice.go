package main

import (
	"fmt"
	"sort"
	"strconv"
)

/**
编写一个程序，提示用户输入整数，并将整数存储在一个已排序的片段中。
程序应编写成一个循环。在进入循环之前，程序应创建一个大小（长度）为 3 的空整数片段。
在每次通过循环时，程序都会提示用户输入一个要添加到片段中的整数。
程序将整数添加到片段中，对片段进行排序，并按排序顺序打印片段的内容。
切片必须不断增大，以容纳用户决定输入的任何数量的整数。
只有当用户输入字符 "X "而不是整数时，程序才会退出（退出循环）。
*/

func main() {
	// 创建一个大小（长度）为 3 的空整数片段
	intSlice := make([]int, 3)

	for {
		// 提示用户输入整数
		var input string
		// 提示用户输入浮点数
		fmt.Print("please enter int: ")
		_, err := fmt.Scanf("%s", &input)
		if err != nil {
			continue
		}

		if input == "X" {
			break
		}

		// 将字符串转成整数
		intInput, _ := strconv.Atoi(input)
		// 将整数添加到片段中
		intSlice = append(intSlice, intInput)

		// 对片段进行排序
		sort.Ints(intSlice)

		fmt.Println(intSlice)
	}
}
