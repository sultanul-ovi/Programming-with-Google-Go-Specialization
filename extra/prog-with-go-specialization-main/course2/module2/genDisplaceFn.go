package main

import (
	"fmt"
)

func main() {
	var a, v0, s0, t float64

	// Prompt user for acceleration, initial velocity, and initial displacement
	fmt.Print("Enter acceleration (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter initial velocity (vo): ")
	fmt.Scan(&v0)
	fmt.Print("Enter initial displacement (so): ")
	fmt.Scan(&s0)

	fmt.Printf("Given a =%f, v0=%f, s0=%f \n", a, v0, s0)

	fn := GenDisplaceFn(a, v0, s0)

	fmt.Print("Enter time (t): ")
	fmt.Scan(&t)

	displacement := fn(t)

	fmt.Printf("Displacement after %.2f seconds: %.2f\n", t, displacement)

}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {

	var f = func(t float64) float64 {

		return 0.5*a*t*t + v0*t + s0
	}

	return f
}
