package main

import "testing"

func TestIsSafeDistance(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected bool
	}{
		{
			name:     "unsafe distance",
			num1:     2,
			num2:     7,
			expected: false,
		},
		{
			name:     "safe distance",
			num1:     3,
			num2:     2,
			expected: true,
		},
		{
			name:     "safe distance",
			num1:     3,
			num2:     0,
			expected: true,
		},
		{
			name:     "unsafe distance",
			num1:     3,
			num2:     8,
			expected: false,
		},
		{
			name:     "unsafe distance",
			num1:     16,
			num2:     9,
			expected: false,
		},
	}
	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsSafeDistance(tt.num1, tt.num2)
			if got != tt.expected {
				t.Errorf("IsSafeDistance(%d, %d) = %t; want %t",
					tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}

func TestReportAnalyzer(t *testing.T) {
	tests := []struct {
		level    []int
		expected bool
	}{
		{
			level:    []int{},
			expected: false,
		},
		{
			level:    []int{1},
			expected: false,
		},
		{
			level:    []int{16, 6, 4, 2, 1},
			expected: true,
		},
		{
			level:    []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			level:    []int{1, 2, 7, 8, 9},
			expected: false,
		},

		{
			level:    []int{9, 7, 6, 2, 1},
			expected: false,
		},
		{
			level:    []int{1, 3, 2, 4, 5},
			expected: true,
		},

		{
			level:    []int{8, 6, 4, 4, 1},
			expected: true,
		},
		{
			level:    []int{1, 3, 6, 7, 9},
			expected: true,
		},
		{
			level:    []int{29, 28, 27, 25, 26, 25, 22, 20},
			expected: true,
		},
		{
			level:    []int{48, 46, 47, 49, 51, 54, 56},
			expected: true,
		},
		{
			level:    []int{1, 1, 2, 3, 4, 5},
			expected: true,
		},
		{ // 1 2 3 4 5 5
			level:    []int{1, 2, 3, 4, 5, 5},
			expected: true,
		},
		{ // 5 1 2 3 4 5
			level:    []int{5, 1, 2, 3, 4, 5},
			expected: true,
		},
		{ // 1 4 3 2 1
			level:    []int{1, 4, 3, 2, 1},
			expected: true,
		},

		{ // 1 6 7 8 9
			level:    []int{1, 6, 7, 8, 9},
			expected: true,
		},
		{ // 1 2 3 4 3
			level:    []int{1, 2, 3, 4, 3},
			expected: true,
		},
		{ // 9 8 7 6 7
			level:    []int{9, 8, 7, 6, 7},
			expected: true,
		},
		{ // 7 10 8 10 11
			level:    []int{7, 10, 8, 10, 11},
			expected: true,
		},
	}
	// Run all test cases
	for _, tt := range tests {
		t.Run("ProblemDampener", func(t *testing.T) {
			got := ProblemDampener(tt.level)
			if got != tt.expected {
				t.Errorf("ProblemDampener(%v) = %t; want %t",
					tt.level, got, tt.expected)
			}
		})
	}
}

func TestCompareNumbers(t *testing.T) {
	tests := []struct {
		name     string
		num1     int
		num2     int
		expected string
	}{
		{
			name:     "first number greater than second",
			num1:     3,
			num2:     2,
			expected: "decrease",
		},
		{
			name:     "first number less than second",
			num1:     2,
			num2:     5,
			expected: "increase",
		},
		{
			name:     "equal numbers",
			num1:     4,
			num2:     4,
			expected: "equal",
		},
		{
			name:     "zero and positive number",
			num1:     0,
			num2:     1,
			expected: "increase",
		},
		{
			name:     "negative numbers",
			num1:     -3,
			num2:     -5,
			expected: "decrease",
		},
		{
			name:     "negative and positive numbers",
			num1:     -1,
			num2:     1,
			expected: "increase",
		},
	}
	// Run all test cases
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CompareNumbers(tt.num1, tt.num2)
			if got != tt.expected {
				t.Errorf("CompareNumbers(%d, %d) = %s; want %s",
					tt.num1, tt.num2, got, tt.expected)
			}
		})
	}
}
