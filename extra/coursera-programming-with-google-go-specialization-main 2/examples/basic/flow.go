package basic

//lint:ignore U1000 (example)
func if_flow() {
	x := 5
	if x > 0 {
		// do something
	} else if x < 0 {
		// do another thing
	} else {
		// do something else
	}
}

//lint:ignore U1000 (example)
func for_loop() {
	for x := 0; x < 5; x++ {
		if x == 2 {continue}
		if x == 3 {break}
	}
}

//lint:ignore U1000 (example)
func switch_case_flow() {
	x := 1
	switch x {
	case 1:
		// 1st case
	case 2:
		// 2nd case
	default:
		// default case
	}

	// tagless switch: case contain boolean
	// => run 1st condition that is true
	switch {
	case x > 0:
		// 1st case
	case x < 0:
		// 2nd case
	default:
		// default case
	}
}
