package main

import "fmt"

func main() {
	var n = []int{1, 39, 2, 9, 7, 54, 11}
	// var n = []int{8, 1, 7, 1, 9, 5, 8, 1, 6, 0, 3, 5, 4, 3, 2, 1, 4, 4, 8, 1, 4, 1, 8, 5, 2, 6}
	// OK
	var isDone = false

	for !isDone {
		isDone = true
		var i = 0
		for i < len(n)-1 {
			if n[i] > n[i+1] {
				n[i], n[i+1] = n[i+1], n[i]
				isDone = false
			}
			i++
		}
	}

	fmt.Println("\n", n)
}
