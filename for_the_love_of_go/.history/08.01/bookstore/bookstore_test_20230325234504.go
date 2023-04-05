package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()
	_ = bookstore.Book{
		Title:  "Spark Joy",
		Author: "Marie Kondō",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	b := bookstore.Book{
		Title: "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}
	want := b.Copies - 1
	result := bookstore.Buy(b)
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}

}