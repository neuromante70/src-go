package main

import (
	"fmt"
)
func main() {

b := bookstore.Book{
	Title: "Nicholas Chuckleby",
	Author: "Charles Dickens",
	Copies: 8,
}

fmt.Println(b)
}