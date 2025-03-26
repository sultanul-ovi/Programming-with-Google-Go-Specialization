package basic

//lint:ignore U1000 (example)
func array_example() {
	// array has fixed size
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	x[0] = 0
	for i, v := range x {
		_ = i // index
		_ = v // value
	}
	for i := range 5 {
		_ = x[i] // value
	}
}

//lint:ignore U1000 (example)
func slice_example() {
	// slice is a "window" of an array (variable size, up to whole array)
	arr := [...]string{"a", "b", "c", "d", "e"}  // ... indicates array
	s1 := arr[1:3]
	_ = len(s1)  // 2, size of s1
	_ = cap(s1)  // 4, capacity = start of slice -> end of array
	// slice literals will create a new underlying array
	sli1 := []int {1, 2, 3}  // nothing in [] => this is a slice, not array
	sli2 := make([]int, 10)  // len = cap = 10
	sli3 := make([]int, 10, 15)  // len = 10, cap = 15
	sli3 = append(sli3, 100)  // if beyond cap then there is time penalty
	_, _, _ = sli1, sli2, sli3
}

//lint:ignore U1000 (example)
func map_example() {
	var idMap1 map[string]int = make(map[string]int)
	_ = idMap1["unknown"]  // return zero if key is not present
	idMap2 := map[string]int{"joe": 123, "max": 456}
	_ = len(idMap2)
	idMap2["jane"] = 789
	delete(idMap2, "joe")
	value, exists := idMap2["joe"]  // check existence of key
	_, _ = value, exists
	for k, v := range idMap2 {
		_ = k  // key
		_ = v  // value
	}
}
