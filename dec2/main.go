package main

import (
	utils "../utils"
	"fmt"
	"strings"
	"regexp"
	"strconv"
)

type characterRange struct {
	min, max int
}

func main() {
	lines := utils.ReadLinesFromFile("passwords_and_policies.txt")
	
	meetCount := 0
	for _, line := range lines {
		charRange, character, password := parseLine(line)
		meetsPolicy := passwordMeetsPolicy(character, charRange, password)
		if meetsPolicy {
			meetCount++
		}
	}

	fmt.Printf("%d passwords meet requirements", meetCount)

}

func parseLine(line string) (characterRange, string, string) {
	reg := regexp.MustCompile(`(\d+-\d+) ([a-z]): (\w+)`)
	res := reg.FindStringSubmatch(line)

	countRange := res[1]
	character := res[2]
	password := res[3]

	i := strings.Index(countRange, "-")
	min, _ := strconv.Atoi(countRange[:i])
	max, _ := strconv.Atoi(countRange[i+1:])

	return characterRange{min, max}, character, password

	
}

func passwordMeetsPolicy(requiredLetter string, requiredLetterCount characterRange, password string) bool {
	count := strings.Count(password, requiredLetter)
	return count >= requiredLetterCount.min && count <= requiredLetterCount.max
}