package bookstore

import (
	"errors"
	"fmt"
)

// Book represents information about a book.
type Book struct {
	ID     int
	Title  string
	Author string
	Copies int
	PriceCents	int
	DiscountPercente	int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalog map[int]Book) []Book {
	result := []Book{}
	for _, b := range catalog {
		result = append(result, b)
	}
	return result
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
func GetBook(catalog map[int]Book, ID int) (Book, error) {
	b, ok := catalog[ID]
	if !ok {
		return Book{}, fmt.Errorf("ID %d doesn't exist", ID)
	}
	return b, nil
}
