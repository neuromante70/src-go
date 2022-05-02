package main

import (
	"fmt"
)

func insert(array []int, lenght int, tempArray []int, value int, insertIndex int) {
	for i := 0; i < lenght; i++ {
		if i < insertIndex {
			tempArray[i] = array[i]
		} else {
			tempArray[i+1] = array[i]
		}
	}
	tempArray[insertIndex] = value
}

func main() {
	var scores = []int{90, 70, 50, 80, 60, 85}
	var lenght = len(scores)
	var tempArray = make([]int, lenght+1)

	insert(scores, lenght, tempArray, 75, 2)
	scores = tempArray

	for i := 0; i <= lenght; i++ {
		fmt.Printf("%d ", scores[i])
	}
}
