package main

import "fmt"

/*
Reserved keywords
any	default	func	interface	select
break	defer	go	map	struct
case	else	goto	package	switch
chan	fallthrouogh	if	range	type
const	for	import	return	var
continue				

The Universe Block (from Learning Go)
Theres actually one more block that is a little weird: the universe block. Remember, Go is a small language with only 25 keywords. Whats interesting is that the built-in types (like int and string), constants (like true and false), and functions (like make or close) arenât included in that list. Neither is nil. So, where are they?
Rather than make them keywords, Go considers these predeclared identifiers and defines them in the universe block, which is the block that contains all other blocks.
Because these names are declared in the universe block, it means that they can be shadowed in other scopes. You can see this happen by running the code in this on The Go Playground.

fmt.Println(true)
true := 10
fmt.Println(true)

When you run it, youâll see:
true
10
You must be very careful to never redefine any of the identifiers in the universe block. If you accidentally do so, you will get some very strange behavior. If you are lucky, youâll get compilation failures. If you are not, youâll have a harder time tracking down the source of your problems.

You might think that something this potentially destructive would be caught by linting tools. Amazingly, it isnât. Not even shadow detects shadowing of universe block identifiers.
*/


func main(){
  /*
  LITERALS
  Integer
  0b for binary
  0 or 0o for octal (don't use the first notation, it's confusing)
  0x for hexadecimal
  */

  vBin := 0b1000
  fmt.Println(vBin)   // print 8 base 10
  vOct := 0o10
  fmt.Println(vOct)   // print 8 base 8 or Oct
  vHex := 0o10
  fmt.Println(vHex)   // print 8 base 16 or Hex
  vLongnum := 1

}
