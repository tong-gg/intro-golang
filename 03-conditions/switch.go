package main

import (
	"fmt"
	"time"
)

func main() {
	today := time.Monday

	/*
		Switch-case in Go is auto-break
		if you want to go to the next case
		use keyword "fallthrough"
	*/
	switch today {
	case time.Monday:
		fmt.Println("Today is Monday")
		fallthrough
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("Hello")
	}
}
