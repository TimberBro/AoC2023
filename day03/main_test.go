package main

import (
	"testing"
)

func TestFirstPart(t *testing.T) {
	got := firstPart("test_input.txt")
	expected := 4361

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
