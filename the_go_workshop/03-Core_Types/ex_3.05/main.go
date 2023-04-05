package main

// TEXT: Safely looping over a String

import (
	"fmt"
)

func main() {
	logLevel := "デバッグ"
	for index, runeVal := range logLevel {
		fmt.Println(index, string(runeVal))
	}
}
