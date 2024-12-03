package main

import (
	"reflect"
	"testing"
)

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
