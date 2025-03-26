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

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

type Bird struct{}

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

type Snake struct{}

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

type Zoo struct {
	animals map[string]Animal
}

func NewZoo() *Zoo {
	return &Zoo{animals: make(map[string]Animal)}
}

func (zoo Zoo) AddAnimal(name string, animal Animal) {
	zoo.animals[name] = animal
	fmt.Println("Created it!")
}

func (zoo Zoo) FindAnimal(name string) (Animal, bool) {
	animal, found := zoo.animals[name]
	return animal, found
}

func main() {
	zoo := NewZoo()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())
		parts := strings.Fields(input)

		if len(parts) != 3 {
			fmt.Println("Invalid command. Please use 'newanimal' or 'query'.")
			continue
		}

		command, name, argument := parts[0], parts[1], parts[2]

		switch command {
		case "newanimal":
			var animal Animal
			switch argument {
			case "cow":
				animal = Cow{}
			case "bird":
				animal = Bird{}
			case "snake":
				animal = Snake{}
			default:
				fmt.Println("Invalid animal type. Use 'cow', 'bird', or 'snake'.")
				continue
			}
			zoo.AddAnimal(name, animal)

		case "query":
			animal, exists := zoo.FindAnimal(name)
			if !exists {
				fmt.Println("Animal not found.")
				continue
			}

			switch argument {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			default:
				fmt.Println("Invalid query type. Use 'eat', 'move', or 'speak'.")
			}
		default:
			fmt.Println("Invalid command. Use 'newanimal' or 'query'.")
		}
	}
}
