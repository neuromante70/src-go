package main

import "fmt"

var y = 42
var z = "Shaken, not stirred"
var e = `James said, "Shaken,
 not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(e)
}
