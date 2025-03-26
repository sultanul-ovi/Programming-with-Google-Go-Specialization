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

type Cow struct {
	food       string
	locomotion string
	noise      string
}

func (a Cow) Eat() {
	fmt.Println(a.food)
}

func (a Cow) Move() {
	fmt.Println(a.locomotion)
}

func (a Cow) Speak() {
	fmt.Println(a.noise)
}

type Bird struct {
	food       string
	locomotion string
	noise      string
}

func (a Bird) Eat() {
	fmt.Println(a.food)
}

func (a Bird) Move() {
	fmt.Println(a.locomotion)
}

func (a Bird) Speak() {
	fmt.Println(a.noise)
}

type Snake struct {
	food       string
	locomotion string
	noise      string
}

func (a Snake) Eat() {
	fmt.Println(a.food)
}

func (a Snake) Move() {
	fmt.Println(a.locomotion)
}

func (a Snake) Speak() {
	fmt.Println(a.noise)
}

func main() {
	const errorString = "Please enter a command of the form of either: \n " +
		"\"newanimal [animalsname] [cow|bird|snake] \" (eg. \"newanimal daisy cow\")\n" +
		"\"query [animalsname] [eat|move|speal] \" (eg. \"query daisy move\")\n" +
		" x to exit\n"

	animalMap := map[string]Animal{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		arguments := strings.Split(text, " ")

		// if they've popped in and "x" over here, we exit
		if strings.ToLower(arguments[0]) == "x" {
			break
		}
		// If we don't have two arguments at this point, it's a syntax error
		if len(arguments) < 3 {
			fmt.Println(errorString)
			continue
		}
		fmt.Println("0: ", arguments[0])
		fmt.Println("1: ", arguments[1])
		fmt.Println("2: ", arguments[2])
		c := arguments[2]
		switch arguments[0] {
		case "newanimal":
			fmt.Printf("sevilla : -%s- \n", c)
			switch c {
			case "cow":
				animalMap[arguments[1]] = Cow{
					food:       "grass",
					locomotion: "walk",
					noise:      "moo",
				}
			case "bird":
				fmt.Println("bird?")
				animalMap[arguments[1]] = Bird{
					food:       "worms",
					locomotion: "fly",
					noise:      "peep",
				}
			case "snake":
				animalMap[arguments[1]] = Snake{
					food:       "mice",
					locomotion: "slither",
					noise:      "hsss",
				}
			
			}
			fmt.Println("Created it!")
			
		case "query":
			animal, ok := animalMap[arguments[1]]
			if !ok {
				fmt.Println("Cannot find animal " + arguments[1])
				fmt.Println(errorString)
				continue
			}
			switch arguments[2] {
			case "eat":
				animal.Eat()
			case "move":
				animal.Move()
			case "speak":
				animal.Speak()
			
			}
			
		default:
			fmt.Println(errorString)
			continue
		}

	}

}
