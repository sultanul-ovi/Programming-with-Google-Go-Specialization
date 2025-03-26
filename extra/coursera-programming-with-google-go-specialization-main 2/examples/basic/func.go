package basic

import "fmt"

//lint:ignore U1000 (example)
func sum_two_int(x, y int) int {
	return x + y
}

//lint:ignore U1000 (example)
func plus_minus_int(x int) (int, int) {
	if x < 0 {
		// passed by value
		// => no external modification
		x = -x
	}
	return x, -x
}

//lint:ignore U1000 (example)
func increase_int(x *int) {
	// passed by reference
	// => external modification
	*x = *x + 1
}

//lint:ignore U1000 (example)
func sum_int_array(a *[3]int) int {
	// avoid making copy of array by passing pointer
	sum := 0
	for i := range 3 {
		sum += (*a)[i]
	}
	return sum
}

//lint:ignore U1000 (example)
func sum_int_slice(a []int) int {
	// slice is a pointer => pass by reference
	sum := 0
	for x := range a {
		sum += x
	}
	return sum
}

//lint:ignore U1000 (example)
func sum_int(vals ...int) int {
	// slice is a pointer => pass by reference
	sum := 0
	for x := range vals {
		sum += x
	}
	return sum
}

//lint:ignore U1000 (example)
func variadic_func_example() int {
	vSlice := []int{1, 2, 3, 4}
	return sum_int(vSlice...)
}

//lint:ignore U1000 (example)
func deferred_example() {
	// Args of deferred call are evaluated immediately
	name := "Beta"
	defer fmt.Println("Bye", name)  // Bye Beta
	name = "Alpha"
	fmt.Println("Hello", name)  // Hello Alpha
}

//lint:ignore U1000 (example)
func func_var_example() func(int) int {
	// return a function
	var funcVar func(int) int
	incFn := func(x int) int {
		return x + 1
	}
	funcVar = incFn
	return funcVar
}

//lint:ignore U1000 (example)
func apply_func(fn func(int) int, x int) int {
	// return a function
	return fn(x)
}

//lint:ignore U1000 (example)
func anon_func_example() {
	// return a function
	apply_func(func(x int) int {return x + 1}, 2)
}
