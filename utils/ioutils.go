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

//Find finds string in slice
func Find(slice []string, val string) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
	}
    return -1, false
}

//FindInt finds int in slice
func FindInt(slice []int, val int) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
	}
    return -1, false
}

//AppendIntIfNotPresent appends if not present
func AppendIntIfNotPresent(slice []int, new int) []int {
	contains := false
	for _, item := range slice {
		if item == new {
			contains = true
		}
	}
	if !contains {
		return append(slice, new)
	} 

	return slice
	
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}