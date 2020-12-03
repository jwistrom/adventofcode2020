package utils

import (
	"os"
	"bufio"
	"log"
)

// ReadLinesFromFile read lines from file
func ReadLinesFromFile(fileName string) []string {
	file, err := os.Open(fileName)
	check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		check(err)
		lines = append(lines, line)
	}

	check(scanner.Err())

	return lines
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}