package main

import (
	"aoc2023/utils"
	"testing"
)

func TestDay01FirstPart(t *testing.T) {
	fileByLines := utils.ReadFileByLines("test_input.txt")

	got := 0
	expected := 142
	for _, line := range fileByLines {
		filteredLine := filterSimpleDigits(line)
		got += combineDigits(filteredLine)
	}

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestDay01SecondPart(t *testing.T) {
	fileByLines := utils.ReadFileByLines("test_second_input.txt")

	got := 0
	expected := 281
	for _, line := range fileByLines {
		first := firstNumber(line)
		last := lastNumber(line)
		got += first*10 + last
	}

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
