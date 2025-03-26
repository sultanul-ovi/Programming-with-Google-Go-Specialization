package main

import (
	"encoding/json"
	"fmt"
)

/**
编写一个程序，提示用户首先输入姓名，然后输入地址。程序应创建一个映射，
并分别使用键 "name "和 "address "将姓名和地址添加到映射中。
您的程序应使用 Marshal() 从地图中创建一个 JSON 对象，然后打印该 JSON 对象。
*/

func main() {
	var (
		name    string
		address string
		MAP     = make(map[string]string)
	)
	// 提示用户首先输入姓名
	fmt.Println("enter name:")
	fmt.Scanf("%s", &name)
	// 然后输入地址
	fmt.Println("enter address:")
	fmt.Scanf("%s", &address)
	// 将姓名和地址添加到映射中
	MAP["name"] = name
	MAP["address"] = address
	// 使用 Marshal() 从地图中创建一个 JSON 对象
	data, _ := json.Marshal(MAP)

	fmt.Printf("%s", data)

}
