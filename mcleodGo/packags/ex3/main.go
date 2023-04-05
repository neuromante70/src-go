package main

import "fmt"

var y = 42

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("binary form: %b\n", y)
	fmt.Printf("exadecimal form: %x\n", y)
	fmt.Printf("exadecimal form with 0x: %#x\n", y)
	y = 911
	fmt.Printf("exadecimal form with 0x: %#x\n", y)
	fmt.Printf("ex - ex 0x - bin: %x\t%#x\t%b\n", y, y, y)
}
