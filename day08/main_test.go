package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 6
		got := firstPart("test_input.txt")

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 6
		got := secondPart("test_second_input.txt")

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

}
