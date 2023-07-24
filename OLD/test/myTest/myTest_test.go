package myTest_test

import (
	"myTest"
	"testing"
)

func TestVimPrint(t *testing.T) {
	t.Parallel()

	want := "vim-go"
	got := myTest.VimPrint()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}

}
