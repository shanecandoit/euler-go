package test

import (
	"euler/src"
	"testing"
)

func TestIsPalindromeNumber(t *testing.T) {
	// IsPalindromeNumber(9009)
	values := map[int]bool{}
	values[0] = true
	values[1] = true
	values[11] = true
	values[12] = false
	values[121] = true
	values[122] = false
	values[1221] = true
	values[9009] = true

	for given, want := range values {
		got := src.IsPalindromeNumber(given)
		if got != want {
			t.Fatalf("given %v got %v wanted %v ", given, got, want)
		}
	}
}

func TestEuler4Known(t *testing.T) {
	given := 2
	got := src.Euler4(given)
	want := []int{91, 99, 9009}
	for i, g := range got {
		if g != want[i] {
			t.Fatalf("given %v got %v wanted %v ", given, got, want)
		}
	}

}
