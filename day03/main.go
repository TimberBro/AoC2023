package main

import (
	"aoc2023/utils"
	"log"
	"strconv"
	"time"
	"unicode"
)

func main() {
	start := time.Now()
	log.Printf("First part result: %d\n", firstPart("day03\\input.txt"))

	//log.Printf("Second part result: %d\n", secondPart())
	elapsed := time.Since(start)
	log.Printf("Execution time %s", elapsed)
}

// TODO: parse file to 2D array
func parseLine(array [][]int32, line string) [][]int32 {
	newArrayLine := make([]int32, 0)

	for _, char := range line {
		newArrayLine = append(newArrayLine, char)
	}

	return append(array, newArrayLine)
}

type Cell struct {
	x int
	y int
	v int32
}

// TODO: find all symbols in 2D array (store it like x/y/v)
func findSymbols(array [][]int32) []Cell {
	cells := make([]Cell, 0)

	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array[i]); j++ {
			if !unicode.IsDigit(array[i][j]) && string(array[i][j]) != "." {
				//log.Printf("Symbol %c was found on %d:%d", array[i][j], i, j)
				cells = append(cells, Cell{
					x: i,
					y: j,
					v: array[i][j],
				})
			}
		}
	}

	return cells
}

// TODO: find all numbers adjacent to symbols, collect them
func numberAround(array [][]int32, i int, j int) []Cell {
	numberCells := make([]Cell, 0)

	// I don't found solution how to inline this boilerplate in loop. Golang doesn't have ternary operators :(
	minX := 0
	if i > 0 {
		minX = -1
	}
	minY := 0
	if j > 0 {
		minY = -1
	}
	maxX := len(array)
	if i < maxX {
		maxX = 1
	} else {
		maxX = 0
	}
	maxY := len(array[0])
	if j < maxY {
		maxY = 1
	} else {
		maxY = 0
	}

	for dx := minX; dx <= maxX; dx++ {
		for dy := minY; dy <= maxY; dy++ {
			if dx != 0 || dy != 0 {
				number := findNumber(array, i+dx, j+dy)
				if number != nil {
					numberCells = append(numberCells, *number)
				}
			}
		}
	}

	return numberCells
}

// TODO: helper function. Find whole number and its head
func findNumber(array [][]int32, x int, y int) *Cell {
	if !unicode.IsDigit(array[x][y]) {
		//log.Printf("initial value is not digit")
		return nil
	}

	partials := make([]int32, 0)
	// going to head of the number
	xFirst := x
	yFirst := 0
	for i := y; i >= 0; i-- {
		if unicode.IsDigit(array[x][i]) {
			partials = append([]int32{array[x][i]}, partials...)
		} else {
			yFirst = i + 1
			break
		}
	}

	// going to back of the number
	for i := y + 1; i < len(array[x]); i++ {
		if unicode.IsDigit(array[x][i]) {
			partials = append(partials, array[x][i])
		} else {
			//log.Printf("Char %c is not digit, break", array[x][y])
			break
		}
	}

	//log.Printf("partials: %v", partials)
	//log.Printf("string partials: %v", string(partials))
	atoi, err := strconv.Atoi(string(partials))
	if err != nil {
		return nil
	}

	return &Cell{
		x: xFirst,
		y: yFirst,
		v: int32(atoi),
	}
}

// TODO: Sum collected numbers
func sumOfParts(numberCells []Cell) int32 {
	var unique []Cell

loop:
	for _, v := range numberCells {
		for i, u := range unique {
			if v.x == u.x && v.y == u.y && v.v == u.v {
				unique[i] = v
				continue loop
			}
		}
		unique = append(unique, v)
	}

	var result int32 = 0
	for _, cell := range unique {
		result += cell.v
	}

	return result
}

func firstPart(file string) int {
	array := make([][]int32, 0)
	fileByLines := utils.ReadFileByLines(file)

	result := 0
	for _, line := range fileByLines {
		array = parseLine(array, line)
	}

	//log.Printf("array: %c\n", array)
	symbols := findSymbols(array)
	//log.Printf("Found symbols. %v", symbols)
	//number := findNumber(array, 0, 4)
	//log.Printf("Whole number = %d found with head on coords: %d:%d", number.v, number.x, number.y)

	//numbersAround := numberAround(array, 1, 3)
	//log.Printf("Found numbers around %d:%d - %+v", 0, 3, numbersAround)

	cells := make([]Cell, 0)
	for _, symbolCell := range symbols {
		cells = append(cells, numberAround(array, symbolCell.x, symbolCell.y)...)
	}

	//log.Printf("Total found numbers around symbols: %v\n", cells)
	result = int(sumOfParts(cells))

	return result
}
