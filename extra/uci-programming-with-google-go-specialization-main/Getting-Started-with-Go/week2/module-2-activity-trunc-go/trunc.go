package main

import (
	"fmt"
	"strconv"
)

/**
编写一个程序，提示用户输入浮点数，并打印出所输入浮点数的截断整数。
截断是去掉小数点右边数字的过程。
*/

func main() {
	// 提示用户输入浮点数
	var input string
	// 提示用户输入浮点数
	fmt.Print("请输入一个浮点数: ")
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	// 使用strconv.ParseFloat解析输入的字符串为浮点数
	flt, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("转换为浮点数失败:", err)
		return
	}

	// 将浮点数转换为整数进行截断
	truncated := int(flt)

	// 打印截断后的整数
	fmt.Printf("截断后的整数是: %d\n", truncated)
}
