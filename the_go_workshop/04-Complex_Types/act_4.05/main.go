package main

// Act 4.05: remove element from slice
import (
	"fmt"
)

func indexOf(element string, data []string) int {
	for k, v := range data {
		if element == v {
			return k
		}
	}
	return -1 // not found
}

func main() {
	series := []string{"Good", "Good", "Bad", "Good", "Good"}
	position := indexOf("Bad", series)
	if position == -1 {
		fmt.Println("Argument not found")
	} else {
		fmt.Println("Bad is in position number", position, "of serie")
	}
	newSeries := append(series[:position], series[position+1:]...)
	fmt.Println("The new serie without bad is", newSeries)
}
