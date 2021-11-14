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
