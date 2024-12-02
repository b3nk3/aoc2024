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
			level:    []int{7, 6, 4, 2, 1},
			expected: true,
		},
		{
			level:    []int{16, 6, 4, 2, 1},
			expected: false,
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
			expected: false,
		},

		{
			level:    []int{8, 6, 4, 4, 1},
			expected: false,
		},

		{
			level:    []int{1, 3, 6, 7, 9},
			expected: true,
		},
	}
	// Run all test cases
	for _, tt := range tests {
		t.Run("ReportAnalyzer", func(t *testing.T) {
			got := ReportAnalyzer(tt.level)
			if got != tt.expected {
				t.Errorf("ReportAnalyzer(%v) = %t; want %t",
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
