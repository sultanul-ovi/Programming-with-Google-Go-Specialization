package main

import (
	"fmt"
)

type Animal_struct struct {
	food, locomotion, noise string
}

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow Animal_struct

func (c Cow) Eat() {
	fmt.Println(c.food)
}
func (c Cow) Move() {
	fmt.Println(c.locomotion)
}
func (c Cow) Speak() {
	fmt.Println(c.noise)
}

type Bird Animal_struct

func (c Bird) Eat() {
	fmt.Println(c.food)
}
func (c Bird) Move() {
	fmt.Println(c.locomotion)
}
func (c Bird) Speak() {
	fmt.Println(c.noise)
}

type Snake Animal_struct

func (c Snake) Eat() {
	fmt.Println(c.food)
}
func (c Snake) Move() {
	fmt.Println(c.locomotion)
}
func (c Snake) Speak() {
	fmt.Println(c.noise)
}

func main() {
	var animal Animal
	var cow Cow = Cow{"grass", "walk", "moo"}
	var bird Bird = Bird{"worms", "fly", "peep"}
	var snake Snake = Snake{"mice", "slither", "hsss"}
	animalMap := make(map[string]Animal) // *** How to use interface correctly with map, the interface only requare for the same function, no matter what kind of type the dynamic type is, only the interface can use different type //*** If here we use Animal_struct(the STRUCT) instead of Animal(the INTERFACE), there will be an error because the struct can not use different type like COW BIRD SNAKE we difined as three different types
	animalMap["cow"] = cow
	animalMap["bird"] = bird
	animalMap["snake"] = snake
	// fmt.Println("If you want to creat new animal, please enter like 'newanimal newanimal_name cow'(notice the type of the new animal can only be cow, bird or snake'; If you want to query for info, please enter like 'query animal_name info'(notice there is default animal name cow, bird and snake, and you can only query for info like eat, move or speak))")
	for {
		var command, animal_name, type_or_method string
		fmt.Print(">")
		fmt.Scan(&command, &animal_name, &type_or_method)
		if command == "newanimal" {
			switch type_or_method {
			case "cow":
				animalMap[animal_name] = cow
			case "bird":
				animalMap[animal_name] = bird
			case "snake":
				animalMap[animal_name] = snake
			}
			fmt.Println("Created it!")
		} else if command == "query" {
			animal = animalMap[animal_name]
			switch type_or_method {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			}
		} else {
			fmt.Println("Invalid command. Please start the program again and enter the string start with either 'newanimal' or 'query'.")
		}
	}

}
