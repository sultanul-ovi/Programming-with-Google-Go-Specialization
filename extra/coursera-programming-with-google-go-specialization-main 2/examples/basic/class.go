package basic

import "math"

type MyInt int

func (mi MyInt) Double() int {
	// mi is passed by value (copied)
	return int(mi * 2)
}

//lint:ignore U1000 (example)
func class_example() {
	v := MyInt(3)
	_ = v.Double()
}

type Point struct {
	// start with lowercase letter => public
	x float64
	y float64
}

// start with uppercase letter => public
func (p Point) DistToOrig() float64 {
	// p is passed by value (cannot change internal vals, copy when called)
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func (p *Point) Scale(v float64) {
	// p is passed by reference (pointer receivers)
	p.x = p.x * v
	p.y = p.y * v
}

//lint:ignore U1000 (example)
func class_func_example() {
	var f1 func(Point) float64 = Point.DistToOrig
	var f2 func(*Point, float64) = (*Point).Scale
	_, _ = f1, f2
}
