package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	
}

func fillArray(arr [10]int) [10]int {
	for i := 0; i < len(arr); i++ {
		arr[i = i + 1]
	}
	return arr
}

func opArray(arr [10]int) [10] int {
	for i := 0; i < len(arr); i++ {
		arr[i]
	}
}