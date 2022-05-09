package main

import "fmt"

func merge(fp []int, sp []int) []int {
	var n = make([]int, len(fp)+len(sp))

	var fpIndex = 0
	var spIndex = 0

	var nIndex = 0

	for fpIndex < len(fp) && spIndex < len(sp) {
		if fp[fpIndex] < sp[spIndex] {
			n[nIndex] = fp[fpIndex]
			fpIndex++
		} else if sp[spIndex] < fp[fpIndex] {
			n[nIndex] = sp[spIndex]
			spIndex++
		}

		nIndex++
	}

	for fpIndex < len(fp) {
		n[nIndex] = fp[fpIndex]

		fpIndex++
		nIndex++
	}

	for spIndex < len(sp) {
		n[nIndex] = sp[spIndex]

		spIndex++
		nIndex++
	}

	return n
}

func mergeSort(arr []int) []int {
	if len(arr) == 1 {
		return arr
	}

	var fp = mergeSort(arr[0 : len(arr)/2])
	var sp = mergeSort(arr[len(arr)/2:])

	return merge(fp, sp)

}

func main() {
	var n = []int{-144, -68, 11, 41, 143, -116, -133, 84, 23, -53, 17, 143, -20, 12}
	// var n = []int{8, 1, 7, 1, 9, 5, 8, 1, 6, 0, 3, 5, 4, 3, 2, 1, 4, 4, 8, 1, 4, 1, 8, 5, 2, 6}
	// Doesn´t work with big lists
	fmt.Println("\n", mergeSort(n))
}
