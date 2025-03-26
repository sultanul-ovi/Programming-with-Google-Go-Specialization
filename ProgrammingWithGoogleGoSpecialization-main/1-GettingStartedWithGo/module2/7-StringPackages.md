# String Packages

## Unicode Package

- Runes are divided into many different categories
- provides a set of functions to test categories of runes
- provides a set of cuntions to test categories of runes
- `IsDigit(r rune)`
- `IsSpace (r rune)`
- `IsLetter(r rune)`
- `IsLower(r rune)`
- `IsPunct (r rune)`
- some functions perform conversions
  - `ToUpper(r rune)`
  - `ToLower(r rune)`
- functions to manipulate UTF-8 encoded strings
- string search functions
  - `Compare(a,b)` returns an integer comparing two strings lexicographically. 0 if a==b, -1 if a<b, +1 if a>b
  - `Contains(s, substr)` returns true if `substr` is inside `s`
  - `HasPrefix(s, prefix)` returns true if the string `s` begins with `prefix`
  - `Index(s, substr)` returns the index of the first instance of `substr` in `s`

## String Manipulation

- strings are immutable but modified strings are returned
- `Replace(s, old, new, n)` returnes a copy of the string `s` with the first `n` instances of old replaced by `new`
- `ToLower(s)`
- `ToUpper(s)`
- `TrimSpace(s)` returns a new string with all leading and trailing white space removed

## Strconv Package

- conversions to and from string represetations of basic data types
- `Atoi(s)` converts string to int
- `Itoa(s)` converts int(base10) to string
- `FormatFloat(f, fmt, prec, bitSize)` converts floating point number to a string
- `ParseFloat(s, bitSize)` converts a string to a floating point number
