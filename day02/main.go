package main

import (
	"aoc2023/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("First part result: %d", firstPart())

}

func firstPart() int {
	fileByLines := utils.ReadFileByLines("day02\\input.txt")

	result := 0
	for _, line := range fileByLines {
		game := parseGame(line)
		if validateRules(game) {
			result += game.ID
		}
	}
	return result
}

type Game struct {
	ID   int
	Sets []Set
}

type Set struct {
	Blue  int
	Red   int
	Green int
}

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGame(line string) Game {
	resultGame := Game{
		ID:   0,
		Sets: make([]Set, 0),
	}
	splittedLine := strings.Split(line, ":")
	gameIdPattern := `Game (\d+)`
	gameIdRegex, _ := regexp.Compile(gameIdPattern)
	gameIdSubmatch := gameIdRegex.FindStringSubmatch(splittedLine[0])
	//fmt.Println(gameIdSubmatch)
	resultGame.ID, _ = strconv.Atoi(gameIdSubmatch[1])

	// TODO: extract to function
	splittedSets := strings.Split(splittedLine[1], ";")
	setsPattern := `\b(\d+)\s+(\w+)\b`
	setsRegex, _ := regexp.Compile(setsPattern)

	for _, set := range splittedSets {
		setsSubmatch := setsRegex.FindAllString(set, -1)
		//fmt.Printf("%+v\n", setsSubmatch)
		s := Set{}
		for _, submatch := range setsSubmatch {
			split := strings.Split(submatch, " ")
			switch split[1] {
			case "blue":
				s.Blue, _ = strconv.Atoi(split[0])
			case "red":
				s.Red, _ = strconv.Atoi(split[0])
			case "green":
				s.Green, _ = strconv.Atoi(split[0])
			}
		}
		resultGame.Sets = append(resultGame.Sets, s)
	}

	return resultGame
}

func validateRules(game Game) bool {
	// provided by author
	maxBlue := 14
	maxGreen := 13
	maxRed := 12

	for _, set := range game.Sets {
		// just return false?
		if set.Blue > maxBlue || set.Red > maxRed || set.Green > maxGreen {
			return false
		}
		continue
	}
	return true
}

// parse line like: "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
func parseSet(line string) Set {
	stringSets := strings.Split(line, ";")

	for _, stringSet := range stringSets {
		strings.Split(stringSet, ",")

	}
	return Set{}
}
