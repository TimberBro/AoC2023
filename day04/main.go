package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	log.Printf("First part result: %d\n", firstPart())

	log.Printf("Second part result: %d\n", secondPart())
	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

func firstPart() int {
	fileByLines := utils.ReadFileByLines("day04\\input.txt")

	result := 0
	for _, line := range fileByLines {
		card := parseCard(line)
		//fmt.Printf("Card %d: winning numbers: %v, existing numbers: %v\n", card.Id, card.winningNumbers, card.existingNumbers)
		numbersIntersect := intersection(card.winningNumbers, card.existingNumbers)
		//fmt.Printf("Intersection of numbers: %v\n", numbersIntersect)
		//fmt.Printf("Counting on line: %d\n", i)
		result += countPoints(numbersIntersect)
	}

	return result
}

func secondPart() int {
	fileByLines := utils.ReadFileByLines("day04\\input.txt")

	result := 0

	cards := make([]Card, 0)
	for _, line := range fileByLines {
		card := parseCard(line)

		cards = append(cards, card)
	}

	result = processCards(cards)

	return result
}

// TODO: parse card
func parseCard(line string) Card {
	resultCard := Card{
		Id:              0,
		winningNumbers:  make([]int, 0),
		existingNumbers: make([]int, 0),
	}

	splittedLine := strings.Split(line, ":")
	cardIdSubmatch := strings.Fields(splittedLine[0])
	resultCard.Id, _ = strconv.Atoi(cardIdSubmatch[1])

	parseNumbers(splittedLine[1], &resultCard)

	return resultCard
}

// parse line like "41 48 83 86 17 | 83 86  6 31 17  9 48 53"
func parseNumbers(line string, card *Card) {
	split := strings.Split(line, "|")
	winningStrNumbers := strings.Fields(split[0])
	existingStrNumbers := strings.Fields(split[1])

	for _, number := range winningStrNumbers {
		atoi, err := strconv.Atoi(number)
		if err != nil {
			fmt.Errorf("unable to parse winning number")
		}
		card.winningNumbers = append(card.winningNumbers, atoi)
	}

	for _, number := range existingStrNumbers {
		atoi, err := strconv.Atoi(number)
		if err != nil {
			fmt.Errorf("unable to parse winning number")
		}
		card.existingNumbers = append(card.existingNumbers, atoi)
	}
}

// TODO: find intersections between two sides
func intersection(slice1 []int, slice2 []int) []int {
	commonValues := make([]int, 0)

	for _, number1 := range slice1 {
		for _, number2 := range slice2 {
			if number1 == number2 {
				commonValues = append(commonValues, number1)
			}
		}
	}

	return commonValues
}

// TODO: count points
func countPoints(slice []int) int {
	// if empty, return 0
	if len(slice) == 0 {
		return 0
	}

	// count through recursion
	if len(slice) == 1 {
		return 1
	} else {
		return 2 * countPoints(slice[1:])
	}
}

type Card struct {
	Id              int
	winningNumbers  []int
	existingNumbers []int
	score           int
}

func processCards(cards []Card) int {
	cardToWinsMap := make(map[int]int)

	// We have every original card from start
	for i := 0; i < len(cards); i++ {
		cardToWinsMap[i] = 1
	}

	for i, card := range cards {
		wins := len(intersection(card.winningNumbers, card.existingNumbers))

		for j := 0; j < cardToWinsMap[i]; j++ {
			for k := 1; k <= wins; k++ {
				cardToWinsMap[i+k] += 1
			}
		}
	}

	totalCards := countTotalCards(cardToWinsMap)

	return totalCards
}

func countTotalCards(cardsToWins map[int]int) int {
	result := 0

	for i := 0; i < len(cardsToWins); i++ {
		result += cardsToWins[i]
	}

	return result
}
