package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T) {
	t.Parallel()

	_ = bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *testing.T) {
	t.Parallel()

	b_ := bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
	want :=1
	result := books


}