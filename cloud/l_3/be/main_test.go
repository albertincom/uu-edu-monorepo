package main

import "testing"

func TestAdd(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{a: 1, b: 2, expected: 3},
		{a: -1, b: 5, expected: 4},
		{a: 0, b: 0, expected: 0},
		{a: -3, b: -7, expected: -10},
		{a: 100, b: 200, expected: 300},
	}

	for _, tc := range testCases {
		result := Add(tc.a, tc.b)
		if result != tc.expected {
			t.Errorf("Add(%d, %d) = %d; want %d", tc.a, tc.b, result, tc.expected)
		}
	}
}
