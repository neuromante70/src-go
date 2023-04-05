package main

import (
	"fmt"
	"bookstore"
)
func main() {

b := bookstore.Book{
	Title: "Nicholas Chuckleby",
	Author: "Charles Dickens",
	Copies: 8,
}

fmt.Println(b)
}