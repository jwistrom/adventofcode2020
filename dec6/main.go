package main

import (
	"fmt"
	"../utils"
	"strings"
	"part2"
)

func main() {

	lines := utils.ReadLinesFromFile("group_answers.txt")

	groupAnswers := getGroupAnswers(lines)
	//fmt.Printf("Group answers: %s\n", groupAnswers)	

	uniqueGroupAnswers := reduceUniqueAnswers(groupAnswers)
	//fmt.Printf("Reduced group answers: %s\n", uniqueGroupAnswers)

	sumUniqueAnswers := sumUniqueAnswers(uniqueGroupAnswers)
	fmt.Printf("Sum of unique answers: %d\n", sumUniqueAnswers)
}

type groupAnswer struct {
	answers []string
}

func sumUniqueAnswers(uniqueAnswers []groupAnswer) (ret int) {
	for _, answer := range uniqueAnswers {
		ret += len(answer.answers)
	}
	return
}

func getGroupAnswers(lines []string) (ret []groupAnswer) {
	var currentGroupAnswers []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			ret = append(ret, groupAnswer{currentGroupAnswers})
			currentGroupAnswers = [] string {}
		}
		for _, char := range line {
			currentGroupAnswers = append(currentGroupAnswers, string(char))
		}
	} 

	ret = append(ret, groupAnswer{currentGroupAnswers})

	return
}

func reduceUniqueAnswers(groupAnswers []groupAnswer) (ret []groupAnswer) {
	for _, answers := range groupAnswers {
			ret = append(ret, answers.uniqueAnswers())
	}
	return
}

func (ga groupAnswer) uniqueAnswers() groupAnswer {
	var ret []string
	for _, answer := range ga.answers {
		_, found := utils.Find(ret, answer)
		if !found {
			ret = append(ret, answer)
		}
	}

	return groupAnswer{ret}
}