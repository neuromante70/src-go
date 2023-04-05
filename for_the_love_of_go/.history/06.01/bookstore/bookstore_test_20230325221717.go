package bookstore_test

import (
	"bookstore"
	//"for_the_love_of_go/06.01/bookstore"
	"testing"
	"errors"
)

// bookstore.Book{
// 	Title:	"Nicholas Chuckleby",
// 	Author:	"Charles Dickens",
// 	Copies:	8
// }

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title:  "Spak Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	b := bookstore.Book{
		Title:  "The return of Sherlock Holmes",
		Author: "Arthur Conan Doyle",
		Copies: 4,
	}
	if b.Copies < 1 {
		return b.errors.New("not enough copy")
	}
	want := b.Copies -1
	result := bookstore.Buy(b)
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}