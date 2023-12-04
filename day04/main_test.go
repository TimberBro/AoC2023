package main

import (
	"aoc2023/utils"
	"testing"
)

func TestDay04FirstPart(t *testing.T) {
	fileByLines := utils.ReadFileByLines("test_input.txt")

	got := 0
	expected := 13
	for _, line := range fileByLines {
		card := parseCard(line)
		//fmt.Printf("Card %d: winning numbers: %v, existing numbers: %v\n", card.Id, card.winningNumbers, card.existingNumbers)
		numbersIntersect := intersection(card.winningNumbers, card.existingNumbers)
		//fmt.Printf("Intersection of numbers: %v\n", numbersIntersect)
		got += countPoints(numbersIntersect)
	}

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}

func TestDay04SecondPart(t *testing.T) {
	fileByLines := utils.ReadFileByLines("test_input.txt")

	got := 0
	expected := 30

	cards := make([]Card, 0)
	for _, line := range fileByLines {
		card := parseCard(line)

		cards = append(cards, card)
	}

	got = processCards(cards)

	if got != expected {
		t.Errorf("expected %d but got %d", expected, got)
	}
}
