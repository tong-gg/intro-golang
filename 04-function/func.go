package main

import (
	"fmt"
)

/*
	Function signature

	func name(params type) return type{
		logic...
	}
*/

func info(name, msg string) {
	fmt.Printf("My name is %s\n", name)
	fmt.Printf("Message: %s\n", msg)
}

func today() string {
	return "Friday"
}

func swap(x, y int) (int, int) {
	return y, x
}

func cal(op func(int, int) int) {
	r := op(4, 5)
	fmt.Println("In cal function")
	fmt.Println("result: ", r)
}

func main() {
	info("Gopher", "Programming Language")
	fmt.Println()
	info("Patipan", "A junior computer science")
	fmt.Println()

	day := today()
	fmt.Println(day)

	a, b := swap(10, 20)
	fmt.Println(a, b)

	plus := func(a int, b int) int { return a + b }
	p := plus(1, 2)
	fmt.Println("1 + 2 = ", p)
	fmt.Printf("type: %T\n", plus)
	cal(plus)

	minus := func(a, b int) int { return a - b }
	cal(minus)
}
