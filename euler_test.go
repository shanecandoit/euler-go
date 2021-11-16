package main

import (
	"testing"

	"euler/src"
)

func TestEuler1_Mult3or5(t *testing.T) {
	want := 233168
	got := src.Euler1_Mult3or5(1000)
	if got != want {
		t.Fatalf("got %v wanted %v ", got, want)
	}
}

func TestEuler2(t *testing.T) {
	want := int64(4613732)
	got := src.Euler2(4_000_000)
	if got != want {
		t.Fatalf("got %v wanted %v ", got, want)
	}
}

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
