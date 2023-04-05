package main

import "fmt"

func main() {
	// var n = []int{1, 39, 2, 9, 7, 54, 11}
	var n = []int{-144, -68, 11, 41, 143, -116, -133, 84, 23, -53, 17, 143, -20, 12}
	// OK

	var i = 1
	for i < len(n) {
		var j = i
		for j >= 1 && n[j] < n[j-1] {
			n[j], n[j-1] = n[j-1], n[j]

			j--
		}

		i++
	}

	fmt.Println("\n", n)
}
