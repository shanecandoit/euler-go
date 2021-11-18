package test

import (
	"euler/src"
	"testing"
)

func TestEuler2(t *testing.T) {
	want := int64(4613732)
	got := src.Euler2(4_000_000)
	if got != want {
		t.Fatalf("got %v wanted %v ", got, want)
	}
}
