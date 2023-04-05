package bookstore

// Book represents information about a book.
Type Book struct {
	Title  string
	Author string
	Copies int
}

func Buy(b Book) Book {
	return Book{}
}