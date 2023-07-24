package main

import (
	"crypto/cipher"
	"fmt"
)

func main() {
	// fmt.Print("\033[2J") //clear the screen befor printing the output in the terminal
	var sl []int
	hat := cap(sl)
	for i := 0; i < 1_000_000; i++ {
		sl = append(sl, i)
		c := cap(sl)
		if c != hat {
			fmt.Println(hat, c)
		}
		hat = c
	}
}