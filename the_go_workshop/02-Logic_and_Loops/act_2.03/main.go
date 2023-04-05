package main

import (
	"fmt"
	// "math/rand"
	// "time"
)

// func bubbleSort() {

// }

func main() {
	// Provide seed
	// rand.Seed(time.Now().Unix())

	// declare and initalize a slice of int with random numbers
	// nums := rand.Perm(7)
	nums := []int{0, 8, 9, 6, 8, 6, 0, 2}

	// Print before sort
	fmt.Println("Before:", nums)

	for swapped := true; swapped; {
		swapped = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swapped = true
			}
		}
	}

	fmt.Println("After:", nums)
}
