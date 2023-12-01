package utils

import (
	"bufio"
	"fmt"
	"os"
)

func ReadFileByLines(filePath string) []string {
	input, err := os.Open(filePath)
	defer input.Close()

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(input)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	return fileLines
}
