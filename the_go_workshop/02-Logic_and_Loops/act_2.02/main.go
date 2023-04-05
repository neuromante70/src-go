package main

import (
	"fmt"
)

func main() {
	topWord := ""
	topCount := 0
	words := map[string]int{
		"Gonna": 3,
		"You": 3,
		"Give": 2,
		"Never": 1,
		"Up": 4,
	}
	for key, value := range words {
		//fmt.Println(key, "=", value)
		if value > topCount {
			topWord = key
			topCount = value
		}
	}

	fmt.Println("Most popular word: ", topWord)
	fmt.Println("With a count of: ", topCount)
	
}
