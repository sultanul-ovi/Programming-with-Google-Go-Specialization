package main

import (
	"fmt"
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

var animals = map[string]Animal{
	/*"cow":   Cow{},
	"bird":  Bird{},
	"snake": Snake{},*/
}

func processReq(req string, a Animal) {
	switch req {
	case "eat":
		a.Eat()
	case "move":
		a.Move()
	case "speak":
		a.Speak()
	}
}

func create(typ string) Animal {
	var a Animal
	switch typ {
	case "cow":
		a = &Cow{}
	case "bird":
		a = &Bird{}
	case "snake":
		a = &Snake{}
	}
	return a
}

func main() {
	var cmd, name, param string
	for {
		fmt.Print(">")
		fmt.Scan(&cmd)
		fmt.Scan(&name)
		fmt.Scan(&param)
		switch cmd {
		case "newanimal":
			a := create(param)
			if a != nil {
				animals[name] = a
				fmt.Println("Created it!")
			}
		case "query":
			if a, ok := animals[name]; ok {
				processReq(param, a)
			}
		}
	}
}
