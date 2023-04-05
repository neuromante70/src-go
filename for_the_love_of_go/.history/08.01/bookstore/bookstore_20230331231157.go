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
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(b Book) {
	want := []string{"same", "same", "same"}
	got := []string{"same", "different", "same"}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	return
}
