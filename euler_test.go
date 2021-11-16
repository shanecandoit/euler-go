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
