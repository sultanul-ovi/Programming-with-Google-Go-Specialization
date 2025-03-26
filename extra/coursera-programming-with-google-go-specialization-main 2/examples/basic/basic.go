package basic

//lint:ignore U1000 (example)
func variable_example() {
	var a, b int
	var c int = 1
	var d = 2
	e := 3
	_, _, _, _, _ = a, b, c, d, e
}

//lint:ignore U1000 (example)
func constant_example() {
	const a, b = 1, "Cool"
	const c int = 2
	_, _, _ = a, b, c
	// enum
	const (
		S = iota
		A = iota
		B = iota
		C = iota
	)
	_, _, _, _ = S, A, B, C
}

//lint:ignore U1000 (example)
func pointer_example() {
	var a int = 1
	var ip *int = &a
	var b int = *ip
	_ = b

	ptr := new(int)
	*ptr = 3
}

//lint:ignore U1000 (example)
func type_conversions_example() {
	var x int16 = 1
	var y int32 = int32(x)
	var t float64 = 1.2345e2  // e is base 10
	var z complex128 = complex(2, 3)  // r, i
	_, _, _ = y, t, z
}
