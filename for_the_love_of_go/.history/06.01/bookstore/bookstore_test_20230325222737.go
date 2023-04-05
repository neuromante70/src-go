package bookstore_test

import (
	"bookstore"
	//"for_the_love_of_go/06.01/bookstore"
	"for_the_love_of_go/06.01/bookstore"
	"testing"
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
		Copies: 2,
	}

	//want := b.Copies - 1
	want := 1
	result, err := bookstore.Buy(b)
	if err != nil {
		t.Fatal(err)
	}
	got := result.Copies
	if want != got {
		t.Errorf("want %d, got %d", want, got)
	}
}

func TestBuyErrorsIfNoCopiesLeft(t *testing.T) {
	t.Parallel()
	b := bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 0,
		}
	_, err := bookstore.Buy(b)
	if err == nil {
		t.Error()
	}
}
