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

func GetBook(catalog []Book, ) []Book {
	return catalog
}