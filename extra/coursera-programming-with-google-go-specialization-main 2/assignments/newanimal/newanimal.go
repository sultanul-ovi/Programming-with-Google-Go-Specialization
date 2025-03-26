package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct {}

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct {}

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct {}

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

var animals map[string]Animal = make(map[string]Animal)

func processNewAnimalRequest(animal_name, animal_type string) {
	var animal Animal
	switch animal_type {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		fmt.Println("animal type", animal_type, "not found in database")
		return
	}
	fmt.Println("Created it!")
	animals[animal_name] = animal
}

func processQueryRequest(animal_name, info_requested string) {
	animal, exists := animals[animal_name]
	if !exists {
		fmt.Println("animal named", animal_name, "not found in database")
		return
	}
	switch info_requested {
	case "eat":
		animal.Eat()
	case "move":
		animal.Move()
	case "speak":
		animal.Speak()
	default:
		fmt.Println("info", info_requested, "is not a valid query")
	}
}

func processRequest(command_str, arg1, arg2 string) {
	switch command_str {
	case "newanimal":
		processNewAnimalRequest(arg1, arg2)
	case "query":
		processQueryRequest(arg1, arg2)
	default:
		fmt.Println("command", command_str, "is not a valid command")
	}
}

func main() {
	for {
		fmt.Print("> ")
		var command_str, arg1, arg2 string
		fmt.Scan(&command_str, &arg1, &arg2)
		processRequest(command_str, arg1, arg2)
	}
}
