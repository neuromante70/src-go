package myStringLen_test

import (
	"mytypes"
	"testing"
)

func TestMyStringLen(t *testing.T) {
	t.Parallel()
	input := mytypes.MyString("Hello, Gophers!")
	want := 10
	got := input.Len()
	if want != got {
		t.Errorf("%q: want len %d, got %d", input, want, got)
	}
}
