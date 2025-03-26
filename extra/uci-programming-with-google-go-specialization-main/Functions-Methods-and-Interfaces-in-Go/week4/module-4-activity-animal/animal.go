package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow结构体表示牛
type Cow struct {
	name string
}

// Eat方法实现了Cow结构体满足Animal接口的Eat()方法
func (c Cow) Eat() {
	fmt.Println("grass")
}

// Move方法实现了Cow结构体满足Animal接口的Move()方法
func (c Cow) Move() {
	fmt.Println("walk")
}

// Speak方法实现了Cow结构体满足Animal接口的Speak()方法
func (c Cow) Speak() {
	fmt.Println("moo")
}

// Bird结构体表示鸟
type Bird struct {
	name string
}

// Eat方法实现了Bird结构体满足Animal接口的Eat()方法
func (b Bird) Eat() {
	fmt.Println("worms")
}

// Move方法实现了Bird结构体满足Animal接口的Move()方法
func (b Bird) Move() {
	fmt.Println("fly")
}

// Speak方法实现了Bird结构体满足Animal接口的Speak()方法
func (b Bird) Speak() {
	fmt.Println("peep")
}

// Snake结构体表示蛇
type Snake struct {
	name string
}

// Eat方法实现了Snake结构体满足Animal接口的Eat()方法
func (s Snake) Eat() {
	fmt.Println("mice")
}

// Move方法实现了Snake结构体满足Animal接口的Move()方法
func (s Snake) Move() {
	fmt.Println("slither")
}

// Speak方法实现了Snake结构体满足Animal接口的Speak()方法
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := make(map[string]Animal)

	for {
		fmt.Print(">")
		var command, name, actionOrType string
		fmt.Scan(&command, &name, &actionOrType)

		switch command {
		case "newanimal":
			switch actionOrType {
			case "cow":
				animals[name] = Cow{name}
			case "bird":
				animals[name] = Bird{name}
			case "snake":
				animals[name] = Snake{name}
			default:
				fmt.Println("unknown animal type")
			}
			fmt.Println("Created it!")
		case "query":
			animal := animals[name]
			if animal != nil {
				switch actionOrType {
				case "eat":
					animal.Eat()
				case "move":
					animal.Move()
				case "speak":
					animal.Speak()
				}
			}
			fmt.Println("unknown name")
		}
	}
}
