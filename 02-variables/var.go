package main

import "fmt"

// package scope
var name string = "Gopher"
var firstName = "Patipan"

// cannot declare at package scope!
// use var keyword instead
// lastName := "Boonsimma"

/*
	The zero value is:
	0 for numeric type
	false for the boolean type
	"" (the empty string) for strings.
*/

/*
	basic type:

	bool

	string

	int int8 int16 int32 int64
	uint uint8 uint16 uint32 uint64 uintptr

	byte // alias for uint8

	rune // alias for int32
		// represents a Unicode code point

	float32 float 64

	complex64 complex128
*/

func main() {
	var zeroVal string
	lastName := "Boonsimma"
	fmt.Println(firstName + " " + lastName)
	function1()
	fmt.Printf("name: %q\nType: %T\n", zeroVal, zeroVal)

	fmt.Println("---Basic type---")
	s := "Gopher"
	i := 5
	f := 3.5
	b := true
	r := 'ต'

	fmt.Printf("string: %v\n", s)
	fmt.Printf("int: %v\n", i)
	fmt.Printf("float: %v\n", f)
	fmt.Printf("bool: %v\n", b)
	fmt.Printf("rune: %v\n", r)
}

func function1() {
	fmt.Println("In function 1: " + name)
}
