package calculator_test

import (
	"calculator"
	"testing"
	"math"
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
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
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
		{a: 2, b: 2, want: 0},
		{a: 2, b: 1, want: 1},
		{a: 5, b: -4, want: 9},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	// var want float64 = 9
	// got := calculator.Multiply(3, 3)
	// if want != got {
	// 	t.Errorf("want %f, got %f", want, got)
	// }
	type testCase struct {
		a, b, want float64
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: -1, b: -1, want: 1},
		{a: 5, b: 0, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, b, want float64
	}

	testCases := []testCase{
		{a: 2, b: 1, want: 2},
		{a: -1, b: -1, want: 1},
		{a: 10, b: 2, want: 5},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		if err != nil {
			t.Fatalf("want no error for valid input, got %v", err)
		}
		if tc.want != got {
			t.Errorf("Divide(%f, %f): want %f, got %f",
				tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("want error for invalid input, got nil")
	}
}

func closeEnough(a, b float64) bool {
	return math.Abs(a-b) <=  0.001
}

func TestSqrt(t *testing.T) {
	t.Parallel()

	type testCase struct {
		a, want float64
	}

	testCases := []testCase{
		{a: 9, want: 3},
		{a: 16, want: 4},
		{a: 98, want: 9.899494937},
	}
	for _, tc := range testCases {
		got := calculator.Sqrt(tc.a)
		// result := closeEnough(tc.want, got)
		// if !result {
		
			t.Errorf("Sqrt(%f): want %f, got %f",
				tc.a, tc.want, got)
		}
	}
}
