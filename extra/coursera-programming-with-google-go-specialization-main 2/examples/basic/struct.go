package basic

import "encoding/json"

type Person struct {
	name	string
	address string
	phone	string
}

//lint:ignore U1000 (example)
func struct_example() {
	p1 := new(Person)
	p2 := Person{name: "joe", address: "SG", phone: "123"}
	p2.name = "max"
	_, _ = p1, p2
	// can also be declared for 1 time usage
	dog := struct {
		name	string
		isGood	bool
	}{"Rex", true}
	_ = dog
}

//lint:ignore U1000 (example)
func json_example() {
	var p1, p2 Person
	p1 = Person{name: "joe", address: "SG", phone: "123"}
	//lint:ignore SA4006,SA9005 (example)
	barr, err := json.Marshal(p1)  // p1 -> barr
	//lint:ignore SA9005 (example)
	err = json.Unmarshal(barr, &p2)  // barr -> p2
	_ = err
}
