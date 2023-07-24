package main

import (
	"fmt"
)

func main() {

	var b byte
	for b = 250; b < 256; b++ {
		fmt.Printf("%d %c\n", b, b)
	}
}
