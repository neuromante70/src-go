package main

import "fmt"

func main() {
	a, b := 5, 10
	fmt.Println("Intial value of a:", a, "and initial value of b:", b)
	// call swap here
	swap(&a, &b)
	fmt.Println(a == 10, b == 5)
	fmt.Println("a is", a, "and b is", b)
}

func swap(a *int, b *int) {
	// swap the values here
	// temp := *a
	// *a = *b
	// *b = temp
	*a, *b = *b, *a
}
