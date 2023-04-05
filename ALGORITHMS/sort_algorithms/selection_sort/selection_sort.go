package main

import "fmt"

func main() {

	sample := []int{-144, -68, 11, 41, 143, -116, -133, 84, 23, -53, 17, 143, -20, 12}
	selectionSort(sample)
}

func selectionSort(arr []int) {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		minIndex := i
		for j := i + 1; j < len; j++ {
			if arr[j] > arr[minIndex] {
				arr[j], arr[minIndex] = arr[minIndex], arr[j]
			}
		}
	}
	fmt.Println("\n", arr)
	// OK
}
