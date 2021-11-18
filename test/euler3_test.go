package test

import (
	"euler/src"
	"testing"
)

func TestEuler3Known(t *testing.T) {
	want := int64(29)
	got := src.Euler3(13195)
	if got != want {
		t.Fatalf("got %v wanted %v ", got, want)
	}
}

func TestEuler3Unknown(t *testing.T) {
	want := int64(6857)
	got := src.Euler3(600851475143)
	if got != want {
		t.Fatalf("got %v wanted %v ", got, want)
	}
}

func TestIsPrime(t *testing.T) {
	values := map[int]bool{}
	values[1] = true
	values[2] = true
	values[3] = true
	values[4] = false
	values[5] = true
	values[6] = false

	for given, want := range values {
		got := src.IsPrime(int64(given))
		if got != want {
			t.Fatalf("given %v got %v wanted %v ", given, got, want)
		}
	}
}
