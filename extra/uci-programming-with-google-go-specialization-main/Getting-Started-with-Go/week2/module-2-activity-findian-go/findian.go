package main

import (
	"fmt"
	"strings"
)

/**
编写一个程序，提示用户输入一个字符串。
程序在输入的字符串中搜索字符 "i"、"a "和 "n"。
如果输入的字符串以字符 "i "开头，以字符 "n "结尾，并包含字符 "a"，则程序应打印 "找到！"。
否则，程序应打印 "未找到！"。程序不区分大小写，因此字符是大写还是小写并不重要。

举例说明 对于以下输入字符串示例，
程序应打印 "找到！"："ian"、"Ian"、"iuiygaygn"、"I d skd a efju N"。
对于以下字符串，程序应打印 "未找到！"："ihhhhhhn"、"ina"、"xian"。
*/

func main() {
	// 提示用户输入一个字符串
	var input string
	fmt.Print("enter string: ")
	_, err := fmt.Scanf("%s", &input)
	if err != nil {
		fmt.Println("enter failed")
		return
	}

	// 去除空格
	input = strings.ReplaceAll(input, " ", "")

	// 转换为小写以不区分大小写
	input = strings.ToLower(input)

	// 检查是否以'i'开头，包含'a'，以及以'n'结尾
	if strings.HasPrefix(input, "i") && strings.Contains(input, "a") && strings.HasSuffix(input, "n") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
