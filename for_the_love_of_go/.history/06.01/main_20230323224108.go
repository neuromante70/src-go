package main

import "fmt"

func foo() {
	fmt.Println("test with ")
}

func main() {
	var title, author string
	var copies, edition, discountPerc int
	var inStock, specialOffer bool
	inStock = true
	specialOffer = false
	var royaltyPercentage float64
	royaltyPercentage = 12.5
	title = "For the Love of go"
	author = "John Arundel"
	copies = 99
	edition = 1
	foo()
	fmt.Println(author, "-", title , "edition Nr:", edition)
	fmt.Println(copies)
}