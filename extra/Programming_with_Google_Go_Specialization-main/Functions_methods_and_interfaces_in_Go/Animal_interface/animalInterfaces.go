package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct{}
type Bird struct{}
type Snake struct{}

func (c Cow) Eat() {
	fmt.Println("grass")
}
func (c Cow) Move() {
	fmt.Println("walk")
}
func (c Cow) Speak() {
	fmt.Println("moo")
}
func (b Bird) Eat() {
	fmt.Println("worms")
}
func (b Bird) Move() {
	fmt.Println("fly")
}
func (b Bird) Speak() {
	fmt.Println("peep")
}
func (s Snake) Eat() {
	fmt.Println("mice")
}
func (s Snake) Move() {
	fmt.Println("slither")
}
func (s Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	fmt.Println("Enter your input in the form <newanimal> <name of new animal> <cow, bird, or snake> /n or<query><animal name><eat, move, or speak> ")
	animals := make(map[string]Animal)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println(">")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if input == "exit" {
			fmt.Println("bye")
			break
		}
		parts := strings.Fields(input)
		if len(parts) != 3 {
			fmt.Println("Invalid input")
			continue
		}
		switch parts[0] {
		case "newanimal":
			name := parts[1]
			animalType := parts[2]
			var newAnimal Animal
			switch animalType {
			case "cow":
				newAnimal = Cow{}
			case "bird":
				newAnimal = Bird{}
			case "snake":
				newAnimal = Snake{}
			default:
				fmt.Println("Invalid type")
				continue
			}
			animals[name] = newAnimal
			fmt.Println("Created it!")
		case "query":
			name := parts[1]
			action := parts[2]
			animal, exists := animals[name]
			if !exists {
				fmt.Println("Animal not found")
				continue
			}
			switch action {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:

				fmt.Println("Invalid action")
			}
		default:
			fmt.Println("Invalid command")
		}

	}
}
