package main

import (
	"fmt"
)

var prt = fmt.Println

func main() {

	var variant interface{}
	variant = "string"
	prt(variant)
	variant = 123
	

}