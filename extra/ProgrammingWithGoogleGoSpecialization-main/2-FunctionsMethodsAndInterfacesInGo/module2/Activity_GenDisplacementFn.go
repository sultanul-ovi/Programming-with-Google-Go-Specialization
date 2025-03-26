package main

import (
	"fmt"
)

func main() {
	var a, v0, s0, t float64
	userInput("Enter acceleration: ", &a)
	userInput("Enter initial velocity: ", &v0)
	userInput("Enter initial displacement: ", &s0)
	userInput("Enter time: ", &t)

	fn := GenDisplaceFn(a, v0, s0)
	fmt.Printf("Displacement after %f time is %f", t, fn(t))
}

func GenDisplaceFn(acceleration float64, initVel float64, initDisplacement float64) func(float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * (time * time)) + (initVel * time) + initDisplacement
	}
}

func userInput(prompt string, time *float64) {
	fmt.Print(prompt)
	_, err := fmt.Scanf("%f", time)
	if err != nil {
		fmt.Println(err)
	}
}
