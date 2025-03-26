package main

import (
	"fmt"
	"os"
	"strings"
)

/**
编写一个程序，从文件中读取信息并用结构体表示。假设有一个包含一系列姓名的文本文件。
文本文件的每一行都有一个名字和一个姓氏，按顺序排列，每行之间用一个空格隔开。

您的程序将定义一个名称结构，该结构有两个字段，fname 代表名字，lname 代表姓氏。
每个字段都是大小为 20 个字符的字符串。

程序应提示用户输入文本文件的名称。程序将连续读取文本文件中的每一行，并创建一个包含文件中名字和姓氏的结构体。
创建的每个结构体都将被添加到一个片段中，在读取完文件中的所有行后，程序中的片段将包含文件中每一行的一个结构体。
读取完文件中的所有行后，您的程序应遍历您的结构体片段，并打印出在每个结构体中找到的第一个和最后一个名称。
*/

type Name struct {
	Fname string
	Lname string
}

func main() {
	// 提示用户输入文件名
	var filename string
	fmt.Print("input file name: ")
	fmt.Scan(&filename)

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("file open failed:", err)
		return
	}
	defer file.Close()

	// 创建一个 Name 结构体的切片
	var names []Name

	// 获取文件信息以确定文件大小
	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println("get file info failed:", err)
		return
	}

	// 读取文件内容到一个字节切片
	fileSize := fileInfo.Size()
	buffer := make([]byte, fileSize)
	_, err = file.Read(buffer)
	if err != nil {
		fmt.Println("read file failed:", err)
		return
	}

	// 将字节切片转换为字符串，并按行分割
	fileContent := string(buffer)
	lines := strings.Split(fileContent, "\n")

	// 处理每一行
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			fname := parts[0]
			lname := parts[1]
			// 将名字和姓氏存入结构体并添加到切片中
			names = append(names, Name{Fname: fname, Lname: lname})
		}
	}

	// 遍历结构体切片并打印每个结构体中的名字和姓氏
	for _, name := range names {
		fmt.Printf("fname: %s, lname: %s\n", name.Fname, name.Lname)
	}
}
