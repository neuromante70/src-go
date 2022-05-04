package main

import "fmt"

func main() {

    sample := []int{8, 1, 7, 1, 9, 5, 8, 1, 6, 0, 3, 5, 4, 3, 2, 1, 4, 4, 8, 1, 4, 1, 8, 5, 2, 6}
    selectionSort(sample)
}

func selectionSort(arr []int) {
    len := len(arr)
    for i := 0; i < len-1; i++ {
        minIndex := i
        for j := i + 1; j < len; j++ {
            if arr[j] < arr[minIndex] {
                arr[j], arr[minIndex] = arr[minIndex], arr[j]
            }
        }
    }
    fmt.Println("\nAfter SelectionSort")
    for _, val := range arr {
        fmt.Println(val)
    }
}