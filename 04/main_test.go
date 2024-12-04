package main

import "testing"

func TestFindCrossMmas(t *testing.T) {
	matrix := [][]string{
		{"M", ".", "S"},
		{".", "A", "."},
		{"M", ".", "S"},
	}

	matrix2 := [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	}

	t.Run("finds XMAS", func(t *testing.T) {
		want := 1
		got, _ := findCrossMas(matrix)
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
	t.Run("finds XMAS", func(t *testing.T) {
		want := 9
		got, _ := findCrossMas(matrix2)
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}
func TestFindXmas(t *testing.T) {
	matrix := [][]string{
		{".", ".", "X", ".", ".", "."},
		{".", "S", "A", "M", "X", "."},
		{".", "A", ".", ".", "A", "."},
		{"X", "M", "A", "S", ".", "S"},
		{".", "X", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", "."},
	}

	matrix2 := [][]string{
		{"M", "M", "M", "S", "X", "X", "M", "A", "S", "M"},
		{"M", "S", "A", "M", "X", "M", "S", "M", "S", "A"},
		{"A", "M", "X", "S", "X", "M", "A", "A", "M", "M"},
		{"M", "S", "A", "M", "A", "S", "M", "S", "M", "X"},
		{"X", "M", "A", "S", "A", "M", "X", "A", "M", "M"},
		{"X", "X", "A", "M", "M", "X", "X", "A", "M", "A"},
		{"S", "M", "S", "M", "S", "A", "S", "X", "S", "S"},
		{"S", "A", "X", "A", "M", "A", "S", "A", "A", "A"},
		{"M", "A", "M", "M", "M", "X", "M", "M", "M", "M"},
		{"M", "X", "M", "X", "A", "X", "M", "A", "S", "X"},
	}
	t.Run("missing matrix", func(t *testing.T) {
		_, err := findXmas(matrix)
		if err != nil {
			t.Errorf("Expected nil, got %v", err)
		}
	})

	t.Run("finds XMAS", func(t *testing.T) {
		want := 4
		got, _ := findXmas(matrix)
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
	t.Run("finds XMAS", func(t *testing.T) {
		want := 18
		got, _ := findXmas(matrix2)
		if got != want {
			t.Errorf("Expected %d, got %d", want, got)
		}
	})
}
