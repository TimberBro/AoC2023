package main

import (
	"aoc2023/utils"
	"log"
	"testing"
)

func TestFirstPart(t *testing.T) {
	array := make([][]int32, 0)
	fileByLines := utils.ReadFileByLines("test_input.txt")

	got := 0
	expected := 4361
	for _, line := range fileByLines {
		array = parseLine(array, line)
	}

	//log.Printf("array: %c\n", array)
	findSymbols(array)
	//number := findNumber(array, 2, 3)
	//log.Printf("Whole number = %d found with head on coords: %d:%d", number.v, number.x, number.y)

	numbersAround := numberAround(array, 0, 3)
	log.Printf("Found numbers around %d:%d - %v", 0, 3, numbersAround)

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
