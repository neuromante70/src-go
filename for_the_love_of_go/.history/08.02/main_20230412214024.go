package main

import "fmt"

type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

func GetBook(catalog map[int]Book, ID int) Book {
	return catalog[ID]
}

func main() {
	catalog := map[int]Book{
		1: {ID: 1, Title: "For the Love of Go"},
		2: {ID: 2, Title: "The Power of Go: Tools"},
	}

	catalog[3] = Book{ID: 3, Title: "Spark Joy"}
	fmt.Println(catalog[3].Title)

	b := catalog[1]
	b.Title = "For the Love of Go esticazzi!"
	catalog[1] = b
	
	fmt.Println(b)

	c, ok := catalog[3]
	if !ok {
		fmt.Println(ok)}
} else {
	
}
