package basic

import "strconv"
import "strings"
import "unicode"

//lint:ignore U1000 (example)
func unicode_example() {
	// unicode rune takes 4 bytes (int32)
	var x rune = '৩'
	_ = unicode.IsDigit(x)
	_ = unicode.IsLower(x)
	_ = unicode.ToLower(x)
}

//lint:ignore U1000 (example)
func string_example() {
	// ASCII char byte takes 1 byte (uint8)
	var x byte = 'a'
	_ = strings.IndexByte("Hanoi", x)
	_ = strings.ToLower("Hanoi")
	_ = strings.Compare("Hanoi", string(x))
	_ = strings.Contains("Hanoi", "an")
	_ = strings.TrimSpace(" Hanoi ")
	_ = strings.ReplaceAll("[a,b,c]", ",", ";")
	_ = strings.Split("Split the words", " ")
	// limiting length of string
	str := "askdfjalsdflsldkfjsl"
	if len(str) > 10 {
		str = str[:10]
	}
	_ = str
}

//lint:ignore U1000 (example)
func string_conversion_example() {
	var x int = 1
	var y string = strconv.Itoa(x)
	x, _ = strconv.Atoi(y)
	_ = x
}
