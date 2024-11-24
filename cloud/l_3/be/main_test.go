package main

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{5, -3, 2},
		{0, 0, 0},
		{-1, -1, -2},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{0, 5, 0},
		{-1, 4, -4},
		{-2, -3, 6},
	}

	for _, tc := range testCases {
		result := Multiply(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Multiply(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}

func TestDivide(t *testing.T) {
	testCases := []struct {
		a, b       int
		expected   int
		shouldFail bool
	}{
		{6, 2, 3, false},
		{5, 0, 0, true}, // Division by zero case
		{-10, -2, 5, false},
	}

	for _, tc := range testCases {
		result, err := Divide(tc.a, tc.b)
		if tc.shouldFail && err == nil {
			t.Errorf("Divide(%d, %d) expected an error but got none", tc.a, tc.b)
		} else if !tc.shouldFail && result != tc.expected {
			t.Errorf("Divide(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
