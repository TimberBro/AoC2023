package main

import (
	"aoc2023/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	// first part
	fileByLines := utils.ReadFileByLines("day01\\input.txt")
	got := 0
	for _, line := range fileByLines {
		filteredLine := filterSimpleDigits(line)
		got += combineDigits(filteredLine)
	}
	fmt.Printf("Answer for first part: %d\n", got)

	// second part
	fileByLines = utils.ReadFileByLines("day01\\input.txt")
	got = 0
	for _, line := range fileByLines {
		first := firstNumber(line)
		last := lastNumber(line)
		got += first*10 + last
	}
	fmt.Printf("Answer for second part: %d\n", got)
}

func filterSimpleDigits(line string) []string {
	re, _ := regexp.Compile("[1-9]")
	digits := re.FindAllString(line, -1)
	return digits
}

func combineDigits(filteredLine []string) int {
	firstDigit := filteredLine[0]
	lastDigit := filteredLine[len(filteredLine)-1]
	atoi, err := strconv.Atoi(firstDigit + lastDigit)

	if err != nil {
		fmt.Errorf("error parsing concatenated strings to integer")
	}

	return atoi
}

// Unable to solve second part using regex, because golang does not support lookahead/lookbehind to parse overlapping values.
// Bruteforce solution
var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func firstNumber(s string) int {
	acc := ""

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc += string(s[i])

		for i, d := range digits {
			if strings.HasSuffix(acc, d) {
				return i + 1
			}
		}
	}
	log.Fatal("not found")
	return 0
}

func lastNumber(s string) int {
	acc := ""

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= '0' && s[i] <= '9' {
			return int(s[i] - '0')
		}

		acc = string(s[i]) + acc

		for i, d := range digits {
			if strings.HasPrefix(acc, d) {
				return i + 1
			}
		}
	}
	return 0
}
