package main

import (
	"strconv"
	"regexp"
	"fmt"
	"../utils"
	"strings"
	"log"
)

func main() {
	lines := utils.ReadLinesFromFile("passports.txt")

	passports := parsePassports(lines)
	fmt.Printf("There are %d passports \n", len(passports))

	validPassports := filterValidPassports(passports)
	fmt.Printf("There are %d valid passports \n", len(validPassports))

	trueValidPassports := filterTrueValidPassports(passports)
	fmt.Printf("There are %d true valid passports \n", len(trueValidPassports))
	
}

func filterValidPassports(passports []passport) (ret []passport) {
	for _, passport :=  range passports {
		if (passport.hasRequiredFields()) {
			ret = append(ret, passport)
		}
	}
	return
}

func filterTrueValidPassports(passports []passport) (ret []passport) {
	for _, passport :=  range passports {
		if (passport.isValid()) {
			ret = append(ret, passport)
		}
	}
	return
}

func parsePassports(lines []string) (ret []passport) {
	
	current := ""
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			ret = append(ret, passport{strings.TrimSpace(current)})
			current = ""
			continue
		}

		current = current + " " + line
	}

	if len(strings.TrimSpace(current)) > 0 {
		ret = append(ret, passport{strings.TrimSpace(current)})
	}

	return
}

type passport struct {
	line string
}

func (p passport) isValid() bool {
	validHgt := p.hasValidHgt()
	validByr := p.hasValidByr()
	validEcl := p.hasValidEcl()
	validEyr := p.hasValidEyr()
	validHcl := p.hasValidHcl()
	validIyr := p.hasValidIyr()
	validPid := p.hasValidPid()

	fmt.Printf("%t - %t - %t - %t - %t - %t - %t\n", validByr, validEcl, validEyr, validHcl, validHgt, validIyr, validPid)

	return p.hasRequiredFields() &&
		validByr &&
		validEcl &&
		validEyr &&
		validHcl && 
		validHgt &&
		validIyr && 
		validPid
}

func (p passport) hasRequiredFields() bool {
	return strings.Contains(p.line, "byr:") &&
	strings.Contains(p.line, "iyr:") &&
	strings.Contains(p.line, "eyr:") &&
	strings.Contains(p.line, "hgt:") &&
	strings.Contains(p.line, "hcl:") &&
	strings.Contains(p.line, "ecl:") &&
	strings.Contains(p.line, "pid:")
}

var hgtRegexp = regexp.MustCompile(`hgt:((\d{2,3})(in|cm))(\s|$)`)
func (p passport) hasValidHgt() bool {
	matches := hgtRegexp.FindStringSubmatch(p.line)
	if (len(matches) != 5) {
		return false
	}

	unit := matches[3]
	height, err := strconv.Atoi(matches[2])
	check(err)

	if unit == "in" {
		return height >= 59 && height <= 76
	} else if unit == "cm" {
		return height >= 150 && height <= 193
	}

	return false
}

var pidRegexp = regexp.MustCompile(`pid:\d{9}(\s|$)`)
func (p passport) hasValidPid() bool {
	return len(pidRegexp.FindStringSubmatch(p.line)) == 2
}

var eclRegexp = regexp.MustCompile(`ecl:(amb|blu|brn|gry|grn|hzl|oth)(\s|$)`)
func (p passport) hasValidEcl() bool {
	matches := eclRegexp.FindStringSubmatch(p.line)
	return len(matches) == 3
}

var hclRegexp = regexp.MustCompile(`hcl:#[a-f0-9]{6}(\s|$)`)
func (p passport) hasValidHcl() bool {
	return len(hclRegexp.FindStringSubmatch(p.line)) == 2
}

var byrRegexp = regexp.MustCompile(`byr:(\d{4})(\s|$)`)
func (p passport) hasValidByr() bool {
	return p.hasValidYear(byrRegexp, 1920, 2002)
}

var iyrRegexp = regexp.MustCompile(`iyr:(\d{4})(\s|$)`)
func (p passport) hasValidIyr() bool {
	return p.hasValidYear(iyrRegexp, 2010, 2020)
}

var eyrRegexp = regexp.MustCompile(`eyr:(\d{4})(\s|$)`)
func (p passport) hasValidEyr() bool {
	return p.hasValidYear(eyrRegexp, 2020, 2030)
}

func (p passport) hasValidYear(regexp *regexp.Regexp, min int, max int) bool {
	matches := regexp.FindStringSubmatch(p.line)
	if (len(matches) != 3) {
		return false
	}

	year, err := strconv.Atoi(matches[1])
	check(err)

	return year >= min && year <= max
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}