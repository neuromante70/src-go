package main

import "fmt"

type Direction int

const (
	North Direction = iota
	South
	East
	West
)

func (d Direction) String() string {
	return [...]string{"North", "South", "East", "West"}[d]
}

func main() {
	var d Direction = South
	fmt.Print(d)

	switch d {
	case North:
		fmt.Println(" goes up.")
	case South:
		fmt.Println(" goes down.")
	default:
		fmt.Println(" stays put.")
	}
}
