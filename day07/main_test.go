package main

import (
	"testing"
)

func TestPart1(t *testing.T) {

	t.Run("Part 1", func(t *testing.T) {
		expected := 6440
		got := part1("test_input.txt")

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

}

func TestPart2(t *testing.T) {

	t.Run("Part 2", func(t *testing.T) {
		expected := 5905
		got := part2("test-input.txt")

		if got != expected {
			t.Errorf("expected %d but got %d", expected, got)
		}
	})

}
