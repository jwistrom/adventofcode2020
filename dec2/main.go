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
	
	meetCount1 := 0
	for _, line := range lines {
		charRange, character, password := parseLine(line)
		meetsPolicy1 := passwordMeetsPolicy1(character, charRange, password)
		if meetsPolicy1 {
			meetCount1++
		}
	}

	fmt.Printf("%d passwords meet requirements 1\n", meetCount1)


	// Part 2

	meetCount2 := 0
	for _, line := range lines {
		charRange, character, password := parseLine(line)
		meetsPolicy2 := passwordMeetsPolicy2(character, charRange, password)
		if meetsPolicy2 {
			meetCount2++
		}
	}

	fmt.Printf("%d passwords meet requirements 2", meetCount2)

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

func passwordMeetsPolicy1(requiredLetter string, requiredLetterCount characterRange, password string) bool {
	count := strings.Count(password, requiredLetter)
	return count >= requiredLetterCount.min && count <= requiredLetterCount.max
}

func passwordMeetsPolicy2(requiredLetter string, requiredLetterCount characterRange, password string) bool {
	if len(password) < requiredLetterCount.max {
		return false
	}

	minLetter := string(password[requiredLetterCount.min-1])
	maxLetter := string(password[requiredLetterCount.max-1])

	valid := (minLetter == requiredLetter && maxLetter != requiredLetter) || (minLetter != requiredLetter && maxLetter == requiredLetter)

	//fmt.Printf("Password is %s. The range is %d-%d. MinLetter is %s. MaxLetter is %s. Required char is %s. The password validity: %t\n", password, requiredLetterCount.min, requiredLetterCount.max, minLetter, maxLetter, requiredLetter, valid)

	return valid
	
}