package main
// Heap memory usage

import (
	"fmt"
	"runtime"
)

func main() {

	var list []int8
	// var list []int8
	for i := 0; i < 1000000000; i++ {
		list = append(list, 100)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("TotalAlloc (Heap) = %v MiB\n", m.TotalAlloc/1024/1024)
}
