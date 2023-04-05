package bookstore

import 

// Book represents information about a book.
type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) (Book, error) {
	if b.Copies < 1 {
		return Book{}, errors.New("not enough copy")
	}
	b.Copies --
	return b, nil
}