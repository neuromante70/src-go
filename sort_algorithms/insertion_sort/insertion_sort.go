package main

import "fmt"

func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}
	// var n = []int{8, 1, 7, 1, 9, 5, 8, 1, 6, 0, 3, 5, 4, 3, 2, 1, 4, 4, 8, 1, 4, 1, 8, 5, 2, 6}
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
