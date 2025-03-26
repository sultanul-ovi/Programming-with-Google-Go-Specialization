package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

var cow, bird, snake Animal = Animal{"grass", "walk", "moo"},
	Animal{"worms", "fly", "peep"},
	Animal{"mice", "slither", "hsss"}

func (animal Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal Animal) Speak() {
	fmt.Println(animal.noise)
}

func main() {
	var ani, request string
	var animal Animal
	for {
		fmt.Print(">")
		fmt.Scanf("%s %s", &ani, &request)
		switch ani {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		}
		switch request {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		}
	}
}
