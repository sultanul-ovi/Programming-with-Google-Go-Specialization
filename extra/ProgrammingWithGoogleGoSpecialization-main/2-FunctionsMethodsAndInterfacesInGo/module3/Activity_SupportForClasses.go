package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func main() {
	for {
		animalType, action := promptForInput()
		processUserInput(animalType, action)
	}
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

var animals = map[string]Animal{
	"cow":   {"grass", "walk", "moo"},
	"bird":  {"worms", "fly", "peep"},
	"snake": {"mice", "slither", "hsss"},
}

func promptForInput() (string, string) {
	fmt.Print("> ")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 2 {
			animalType, action := strings.ToLower(parts[0]), strings.ToLower(parts[1])
			if isValidAnimalType(animalType) && isValidAction(action) {
				return animalType, action
			}
		}

		fmt.Println("Invalid input. It should be an animal type and an action.")
	}

	return promptForInput()
}

func isValidAnimalType(animalType string) bool {
	if _, ok := animals[animalType]; !ok {
		fmt.Println("Invalid animal type. It should be either cow, bird, snake.")
		return false
	}
	return true
}

func isValidAction(action string) bool {
	switch action {
	case "eat", "move", "speak":
		return true
	default:
		fmt.Println("Invalid action. It should be either eat, move, speak.")
		return false
	}
}

func processUserInput(animalType, action string) {
	animal := animals[animalType]
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
