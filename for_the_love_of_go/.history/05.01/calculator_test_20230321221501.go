package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	// var want float64 = 4
	// got := calculator.Add(2, 2)
	// if want != got {
	// 	t.Errorf("want %f, got %f", want, got)
	// }
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want:4},
		{a: 1, b: 1, want:2},
		{a: 5, b: 0, want:5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	// var want float64 = 2
	// got := calculator.Subtract(4, 2)
	// if want != got {
	// 	t.Errorf("want %f, got %f", want, got)
	// }
	type testCase struct {
		a, b float64
		want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want:4},
		{a: 1, b: 1, want:2},
		{a: 5, b: 0, want:5},
	}	
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	var want float64 = 9
	got := calculator.Multiply(3, 3)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
