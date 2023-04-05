package main

import (
	"fmt"
	"for_the_love_of_go/06.01/bookstore"
)

func main() {
	b := bookstore.Book{
		Title:  "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}
	fmt.Println(b)
}

func Buy(myBook bookstore.Book) Book {
	return myBook{}
}
