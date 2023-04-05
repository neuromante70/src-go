package bookstore

import (
	"bookstore"
	"testing"
)

bookstore.Book{
	Title:	"Nicholas Chuckleby",
	Author:	"Charles Dickens",
	Copies:	8
}

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		
	}
}