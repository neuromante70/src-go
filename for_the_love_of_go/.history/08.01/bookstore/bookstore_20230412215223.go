package bookstore

import (
	"errors"
)

// Book represents information about a book.
type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog []Book) []Book {
	return catalog
}

// version with slice
// func GetBook(catalog []Book, ID int) Book {
// for _, b := range catalog {
// 	if b.ID == ID {
// 		return b
// 	}
// }
// 	return Book{}
// }

// version with map
func GetBook(catalog map[int]Book, ID int) (Book, error ){
	b, ok := catalog
	return catalog[ID], nil
}
