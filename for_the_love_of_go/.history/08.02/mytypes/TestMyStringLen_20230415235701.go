package myStringLen_test

import (
	//"for_the_love_of_go/08.0/mytypes"
	"testing"
)

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("Hello culo")
	want := 10
	got := input.Len()
	if want != got {
		t.Errorf("%q: want len %d, got %d", input, want, got)
	}
}
