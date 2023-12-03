package main

import (
	"aoc2023/utils"
	"testing"
)

func TestDay02FirstPart(t *testing.T) {
	fileByLines := utils.ReadFileByLines("test_input.txt")

	got := 0
	expected := 8
	for _, line := range fileByLines {
		game := parseGame(line)
		if validateRules(game) {
			got += game.ID
		}
	}

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestDay02SecondPart(t *testing.T) {
	fileByLines := utils.ReadFileByLines("test_input.txt")

	got := 0
	expected := 2286
	for _, line := range fileByLines {
		game := parseGame(line)
		got += leastCubesRequired(game)
	}

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
