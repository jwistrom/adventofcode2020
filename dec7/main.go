package main

import (
	"../utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"errors"
)

func main() {

	lines := utils.ReadLinesFromFile("bag_rules.txt")
	rules := parseRules(lines)

	colorsContainingGold := colorsContainingColor("shiny gold", rules)
	fmt.Printf("Color eventually containing shiny gold: %d\n", len(colorsContainingGold))

	allBagsInGold := countAllBagsInside("shiny gold", rules)
	fmt.Printf("There are totally %d bags in gold\n", allBagsInGold)
	
}

func countAllBagsInside(color string, rules []rule) (ret int) {
	rule := findRuleWithColor(color, rules)

	for _, containingQty := range rule.content {
		ret += containingQty.qty
		ret += countAllBagsInside(containingQty.color, rules) * containingQty.qty
	}

	return
}

func findRuleWithColor(color string, rules []rule) rule {
	var found = []rule{}

	for _, rule := range rules {
		if rule.color == color {
			found = append(found, rule)
		}
	}

	if len(found) > 1 {
		panic(errors.New("Too many rules found"))
	}

	if len(found) == 0 {
		panic(errors.New("No rule found"))
	}

	return found[0]
}

func colorsContainingColor(containingColor string, rules []rule) (ret []string) {
	for _, rule := range rules {
		if rule.containsColor(containingColor) {
			ret = appendIfNotPresent(ret, rule.color)
			otherColors := colorsContainingColor(rule.color, rules)
			ret = appendAllIfNotPresent(ret, otherColors)
		}
	}

	return
}

func appendAllIfNotPresent(slice []string, new []string) []string {
	for _, toAdd := range new {
		slice = appendIfNotPresent(slice, toAdd)
	}
	return slice
}

func appendIfNotPresent(slice []string, new string) []string {
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

var regex = regexp.MustCompile(`^(\w+ \w+) bags contain (((\d+) (\w+ \w+) bags?(, |\.))+)$`)
var subRegex = regexp.MustCompile(`(((\d+) (\w+ \w+) bags?)+)`)
var noOtherBagsRegex = regexp.MustCompile(`^(\w+ \w+)`)
func parseRules(lines []string) (ret []rule) {

	for _, line := range lines {
		if strings.Contains(line, "no other bags") {
			color := noOtherBagsRegex.FindStringSubmatch(line)[0]
			ret = append(ret, rule{color, []colorQty{}})
			continue
		}

		matches := regex.FindStringSubmatch(line)
		content := subRegex.FindAllStringSubmatch(matches[2], -1)

		color := matches[1]

		var qtys = []colorQty{}
		for _, cnt := range content {
			qty, _ := strconv.Atoi(cnt[3])
			color := cnt[4]
			qtys = append(qtys, colorQty{qty, color})
		}

		ret = append(ret, rule{color, qtys})
	}

	return
}

type rule struct {
	color string
	content []colorQty
}

type colorQty struct {
	qty int
	color string
}

func (r rule) containsColor(color string) bool {
	for _, cnt := range r.content {
		if cnt.color == color {
			return true
		}
	}
	return false
}