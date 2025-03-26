package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a, v0, s0, t float64
	)
	fmt.Println("Please enter your acceleration:")
	fmt.Scan(&a)
	fmt.Println("Please enter your initial velocity:")
	fmt.Scan(&v0)
	fmt.Println("Please enter your initial displacement:")
	fmt.Scan(&s0)
	Desplace := GenDisplaceFn(a, v0, s0)
	fmt.Println("Please enter your time:")
	fmt.Scan(&t)
	fmt.Println(Desplace(t))
}

func GenDisplaceFn(acc, iniVelo, iniDis float64) func(float64) float64 {
	fn := func(time float64) float64 {
		return acc*math.Sqrt(time)/2 + iniVelo*time + iniDis
	}
	return fn
}
