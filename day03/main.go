package main

import (
	"log"
	"strconv"
	"unicode"
)

func main() {

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
				numberCells = append(numberCells, findNumber(array, i+dx, j+dy))
			}
		}
	}

	return numberCells
}

// TODO: helper function. Find whole number and its head
func findNumber(array [][]int32, x int, y int) Cell {
	//if !unicode.IsDigit(array[x][y]) {
	//	log.Fatalf("initial value is not digit")
	//	return Cell{}
	//}

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

	log.Printf("partials: %v", partials)
	log.Printf("string partials: %v", string(partials))
	atoi, err := strconv.Atoi(string(partials))
	if err != nil {
		return Cell{}
	}

	return Cell{
		x: xFirst,
		y: yFirst,
		v: int32(atoi),
	}
}

// TODO: Sum collected numbers
