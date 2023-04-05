package main

import "fmt"

type Currency float64

func (c Currency) String() string {
	return fmt.Sprintf("$%.2f", float64(c))
}

func main() {
	// declaring an interface

	type Empty interface {
	}
	var e Empty
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	e = 7
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	e = "Hello, World!"
	fmt.Printf("e's value: %v, type: %T\n", e, e)

	f := 3.14
	e = &f
	fmt.Printf("e's value: %v, type: %t\n", e, e)

	e = new(Currency)
	fmt.Printf("e's value: %v, type: %t\n", e, e)
}
