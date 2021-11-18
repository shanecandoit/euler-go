package test

import (
	"euler/src"
	"testing"
)

func TestEuler1_Mult3or5(t *testing.T) {
	want := 233168
	got := src.Euler1_Mult3or5(1000)
	if got != want {
		t.Fatalf("got %v wanted %v ", got, want)
	}
}
