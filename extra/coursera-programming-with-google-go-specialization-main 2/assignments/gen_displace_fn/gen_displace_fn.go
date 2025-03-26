package main

import "fmt"

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return a * t * t / 2 + v0 * t + s0
	}
}

func input_float(input_message string) float64 {
	var x float64
	fmt.Print(input_message)
	fmt.Scan(&x)
	return x
}

func main() {
	a := input_float("Input acceleration: ")
	v0 := input_float("Input initial velocity: ")
	s0 := input_float("Input initial displacement: ")
	fn := GenDisplaceFn(a, v0, s0)
	t := input_float("Input time: ")
	fmt.Println("Calculated displacement:", fn(t))
}
