package main

import(
	"../../utils"
	"strings"
	"fmt"
)

func main() {

	lines := utils.ReadLinesFromFile("../group_answers.txt")

	groupAnswers := getGroupAnswers(lines)
	fmt.Printf("Group answers by person: %s\n", groupAnswers)

	sum := sumNumberOfQuestionsAllAnsweredYesTo(groupAnswers)
	fmt.Printf("Sum of questions all answered yes to: %d\n", sum)

}

type groupAnswer struct {
	answers [][]string
}

func sumNumberOfQuestionsAllAnsweredYesTo(groupAnswers []groupAnswer) (ret int) {
	for _, groupAnswer := range groupAnswers {
		ret += groupAnswer.numberOfQuestionsAllAnsweredYesTo()
	}
	return
}

func (ga groupAnswer) numberOfQuestionsAllAnsweredYesTo() (ret int) {
	if len(ga.answers) == 1 {
		fmt.Printf("Number of questions alla answered ues to: %d\n", len(ga.answers[0]))
		return len(ga.answers[0])
	}

	numberOfPersonsInGroup := len(ga.answers)
	answersOfFirstPerson := ga.answers[0]

	for _, firstPersonAnswer := range answersOfFirstPerson {
		var personCount int
		for _, person := range ga.answers {
			_, found := utils.Find(person, firstPersonAnswer)
			if found {
				personCount++
			}
		}
		if personCount == numberOfPersonsInGroup {
			ret++
		}
	}

	fmt.Printf("Number of questions alla answered ues to: %d\n", ret)
	return
}


func getGroupAnswers(lines []string) (ret []groupAnswer) {
	var currentGroupAnswers [][]string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			ret = append(ret, groupAnswer{currentGroupAnswers})
			currentGroupAnswers = [][] string {}
			continue
		}
		var person []string
		for _, char := range line {
			person = append(person, string(char))
		}

		currentGroupAnswers = append(currentGroupAnswers, person)
	} 

	ret = append(ret, groupAnswer{currentGroupAnswers})

	return
}