package main

import "fmt"

func main() {
	var s string // default to ""
	fmt.Printf("string: %#v\n", s)
	var r rune      // default to 0
	fmt.Printf("rune: %#v\n", r)
	var bt byte     // default to 0
	fmt.Printf("byte: %#v\n", bt)
	var i int       // default to 0
	
	var ui uint     // default to 0
	var f float32   // default to 0
	var c complex64 // default to 0
	var b bool      // default to false
	var arr [2]int  // default to [0 0]
	var obj struct {
		b   bool
		arr [2]int
	} // default to {false [0 0]}
	var si []int          // default to int
	var ch chan string    // default to nil
	var mp map[int]string // default to nil
	var fn func()         // default to nil
	var ptr *string       // default to nil
	var all any           // default to nil

}
