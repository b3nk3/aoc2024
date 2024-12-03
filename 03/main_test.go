package main

import (
	"reflect"
	"testing"
)

func TestFilterOutBlocks(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+don't()mul(32,64](mul(11,8)undo()?mul(8,5))"

	t.Run("test filtering out text", func(t *testing.T) {
		got := FilterOutBlocks(input)
		want := "xmul(2,4)&mul[3,7]!^?mul(8,5))"

		if got != want {
			t.Errorf("SplittingOnMultipleDelimiters() = %v, want %v", got, want)
		}
	})

	// t.Run("FindMulsWithRegexp", func(t *testing.T) {
	// 	filteredInput := FilterOutBlocks(input)
	// 	got := FindMulsWithRegexp(filteredInput)
	// 	want := []string{"mul(2,4)", "mul(8,5)"}

	// 	if !reflect.DeepEqual(got, want) {
	// 		t.Errorf("FindMulsWithRegexp() = %v, want %v", got, want)
	// 	}
	// })
	// t.Run("FindMulsWithRegexp", func(t *testing.T) {
	// 	filteredInput := FilterOutBlocks(input)
	// 	matches := FindMulsWithRegexp(filteredInput)
	// 	got := MultiPlySliceItems(matches)
	// 	want := 48

	// 	if !reflect.DeepEqual(got, want) {
	// 		t.Errorf("FindMulsWithRegexp() = %v, want %v", got, want)
	// 	}
	// })

}

func TestFindMulsWithRegexp(t *testing.T) {
	got := FindMulsWithRegexp("xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))")
	want := []string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("FindMulsWithRegexp() = %v, want %v", got, want)
	}
}

func TestMultiPlySliceItems(t *testing.T) {
	got := MultiPlySliceItems([]string{"mul(2,4)", "mul(5,5)", "mul(11,8)", "mul(8,5)"})
	want := 161

	if got != want {
		t.Errorf("MultiPlySliceItems() = %v, want %v", got, want)
	}
}
