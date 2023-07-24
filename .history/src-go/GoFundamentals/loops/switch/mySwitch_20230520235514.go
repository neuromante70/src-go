package main

import (
	"fmt"
)

var ft = fmt.Println

func main() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			ft(i, "is even")
		case i%3 == 0:
			ft(i, "is divisibile by 3 but not 2")
		case i%7 == 0:
			ft("exit the loop!")
			break loop
		default:
			ft(i, "is boring")
		}
	}
}
