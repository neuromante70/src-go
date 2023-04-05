package bookstore

import (
	"errors"
)

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies < 1 {
		return Book{}, errors.New("not copies left")
	}
	b.Copies--
	return b, nil
}