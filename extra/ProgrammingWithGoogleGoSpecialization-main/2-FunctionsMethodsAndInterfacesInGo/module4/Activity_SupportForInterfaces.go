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

func (c Cow) Eat()   { fmt.Println("grass") }
func (c Cow) Move()  { fmt.Println("walk") }
func (c Cow) Speak() { fmt.Println("moo") }

func (b Bird) Eat()   { fmt.Println("worms") }
func (b Bird) Move()  { fmt.Println("fly") }
func (b Bird) Speak() { fmt.Println("peep") }

func (s Snake) Eat()   { fmt.Println("mice") }
func (s Snake) Move()  { fmt.Println("slither") }
func (s Snake) Speak() { fmt.Println("hsss") }

func main() {
	animals := map[string]Animal{
		"cow":   Cow{},
		"bird":  Bird{},
		"snake": Snake{},
	}

	for {
		animalType, action := promptForInput()
		animal, found := animals[animalType]
		if !found {
			fmt.Println("Invalid animal type. It should be either cow, bird, snake.")
			continue
		}
		processUserInput(animal, action)
	}
}

func promptForInput() (string, string) {
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 2 {
			animalType, action := strings.ToLower(parts[0]), strings.ToLower(parts[1])
			return animalType, action
		}

		fmt.Println("Invalid input. It should be an animal type and an action.")
	}

	return promptForInput()
}

func processUserInput(animal Animal, action string) {
	switch action {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("Invalid action. It should be either eat, move, speak.")
	}
}
