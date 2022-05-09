package main

import "fmt"

func bubbleSort(arr []int, size int) []int {
	if size == 1 {
		return arr
	}

	var i = 0
	for i < size-1 {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}

		i++
	}

	bubbleSort(arr, size-1)

	return arr
}

func main() {
	var n = []int{-144, -68, 11, 41, 143, -116, -133, 84, 23, -53, 17, 143, -20, 12}
	// var n = []int{8, 1, 7, 1, 9, 5, 8, 1, 6, 0, 3, 5, 4, 3, 2, 1, 4, 4, 8, 1, 4, 1, 8, 5, 2, 6}
	// OK

	fmt.Println("\n", bubbleSort(n, len(n)))
}
