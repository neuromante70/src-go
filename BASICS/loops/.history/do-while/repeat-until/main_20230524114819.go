package main

import "fmt"

/*
To write a repeat-until loop

repeat
	work();
until condition;

simply change the condition in the code above to its complement:

1)
for ok := true; ok; ok = !condition {
	work()
}

2)
for {
	work()
	if !condition {
		break
	}
}

*/


func main() {

}