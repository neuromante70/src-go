package main

import (
	"fmt"
)

func main() {

var b byte
for b = 250; b <= 255; b++ {
	fmt.Printf("%d %c\n", b, b)
}
}
