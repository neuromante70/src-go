package main

import "fmt"

var foo string = "bar"

func main() {
	var baz string = "qux"
	foo = "culo"
	fmt.Println(foo, baz)
}