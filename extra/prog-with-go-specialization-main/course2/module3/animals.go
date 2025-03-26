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

func (a *Animal) Eat() string {
	return a.food
}

func (a *Animal) Move() string {
	return a.locomotion

}

func (a *Animal) Speak() string {
	return a.noise

}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hsss "}

	animals := map[string]Animal{"cow": cow, "bird": bird, "snake": snake}

	for {
		fmt.Print(">")

		scanner := bufio.NewScanner(os.Stdin)

		scanner.Scan()
		input := scanner.Text()

		parts := strings.Split(input, " ")

		if len(parts) != 2 {
			fmt.Println("Requests must contain 2 strings. Try again")
			continue
		}

		animal := parts[0]
		command := parts[1]

		if x, found := animals[animal]; found {
			if command == "eat" {
				fmt.Println(x.Eat())
			} else if command == "move" {
				fmt.Println(x.Move())
			} else if command == "speak" {
				fmt.Println(x.Speak())
			} else {
				fmt.Println("Unknown request")
			}
		} else {
			fmt.Printf("Unknown animal: %s . Try with cow, bird, snake", animal)
		}
	}

}
