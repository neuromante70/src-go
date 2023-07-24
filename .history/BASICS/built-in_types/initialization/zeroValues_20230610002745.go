package main

import "fmt"

func main() {
	var s string // default to ""
	fmt.Printf("string: %#v\n", s)
	var r rune // default to 0
	fmt.Printf("rune: %#v\n", r)
	var bt byte // default to 0
	fmt.Printf("byte: %#v\n", bt)
	var i int // default to 0
	fmt.Printf("int: %#v\n", i)
	var ui uint // default to 0
	fmt.Printf("uint: %#v\n", ui)
	var f float32 // default to 0
	fmt.Printf("float32: %#v\n", f)
	var c complex64 // default to 0
	fmt.Printf("complex64: %#v\n", c)
	var b bool // default to false
	fmt.Printf("bool: %#v\n", b)
	var arr [2]int // default to [0 0]
	fmt.Printf("array [2]int: %#v\n", arr)
	var obj struct {
		b   bool
		arr [2]int
	} // default to {false [0 0]}
	fmt.Printf("struct: %#v\n", obj)
	var si []int // default to int
	fmt.Printf("slice []int: %#v\n", si)
	var ch chan string    // default to nil
	fmt.Printf("channel chan string: %#v\n",arr)
	var mp map[int]string // default to nil
	fmt.Printf("map map: %#v\n",arr)
	var fn func()         // default to nil
	var ptr *string       // default to nil
	var all any           // default to nil

}
