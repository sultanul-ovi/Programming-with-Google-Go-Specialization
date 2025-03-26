package basic

import "math"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

//lint:ignore U1000 (example)
type Point2D struct {
	x, y float64
}

//lint:ignore U1000 (example)
type Triangle struct {
	a, b, c Point2D
}

func (t Triangle) Area() float64 {
	ab := Point2D{t.b.x - t.a.x, t.b.y - t.a.y}
	ac := Point2D{t.c.x - t.a.x, t.c.y - t.a.y}
	result := math.Abs(ab.x * ac.y - ab.y * ac.x) / 2.0
	return result
}

func (t Triangle) Perimeter() float64 {
	ab := Point2D{t.b.x - t.a.x, t.b.y - t.a.y}
	ac := Point2D{t.c.x - t.a.x, t.c.y - t.a.y}
	bc := Point2D{t.c.x - t.b.x, t.c.y - t.b.y}
	result := 0.0
	result += math.Sqrt(ab.x * ab.x + ab.y * ab.y)
	result += math.Sqrt(ac.x * ac.x + ac.y * ac.y)
	result += math.Sqrt(bc.x * bc.x + bc.y * bc.y)
	return result
}

//lint:ignore U1000 (example)
type Rectangle struct {
	a, b, c, d Point2D
}

func (r Rectangle) Area() float64 {
	ab := Point2D{r.b.x - r.a.x, r.b.y - r.a.y}
	ad := Point2D{r.d.x - r.a.x, r.d.y - r.a.y}
	result := math.Abs(ab.x * ad.y - ab.y * ad.x)
	return result
}

func (r Rectangle) Perimeter() float64 {
	ab := Point2D{r.b.x - r.a.x, r.b.y - r.a.y}
	ad := Point2D{r.d.x - r.a.x, r.d.y - r.a.y}
	result := 0.0
	result += math.Sqrt(ab.x * ab.x + ab.y * ab.y) * 2.0
	result += math.Sqrt(ad.x * ad.x + ad.y * ad.y) * 2.0
	return result
}

//lint:ignore U1000 (example)
func interface_example() {
	t1 := Triangle{Point2D{0, 0}, Point2D{0, 1}, Point2D{1, 0}}
	var t2 *Triangle
	var shape Shape2D  // No dynamic type
	shape = t1  // Triangle is dynamic type, t1 is dynamic value
	shape = t2  // Triangle is dynamic type, but no dynamic value
	_ = shape
}

//lint:ignore U1000 (example)
func interface_func_example(s Shape2D) bool {
	return (s.Area() <= 100 && s.Perimeter() <= 100)
}

//lint:ignore U1000 (example)
func type_assertion_example(s Shape2D) {
	// Type cast
	tri, ok := s.(Triangle)
	if ok {
		// s is Triangle
		_ = tri
	}
	// Type switch
	switch sh := s.(type) {
	case Triangle:
		// sh is Triangle
		_ = sh
	case Rectangle:
		// sh is Rectangle
	}
}
