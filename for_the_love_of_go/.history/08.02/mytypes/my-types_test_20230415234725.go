package mytypes_test

import (
	"mytypes"
	"testing"
	??"the_go_workshop/Go_books/For_the_Love_of_Go/for_the_love_of_go-code/08.02/mytypes"
)

func TestTwice(t *testing.T) {
	t.Parallel()
	input := mytypes.MyInt(9)
	want := mytypes.MyInt(18
	got := input.Twice()
	if want != got {
		t.Errorf("twice %d: want %d, got %d", input, want, got)
	}
}
