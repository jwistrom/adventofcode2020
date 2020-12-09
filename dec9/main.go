package main

import (
	"fmt"
	"../utils"
	"strconv"
	"math"
)

const preamble int = 25
const fileName string = "numbers.txt"

func main() {

	numbers := parseInput()

	faultyNumber := findFaultyNumber(numbers)
	fmt.Printf("The faulty number is %d\n", faultyNumber)

	contiguousSet := findContiguousSet(numbers, faultyNumber)
	fmt.Println(contiguousSet)

	min := min(contiguousSet)
	max := max(contiguousSet)

	fmt.Printf("Sum of smallest (%d) and biggest (%d) in contiguous set is %d\n", min, max, min + max)

}

func findContiguousSet(numbers []int, targetSum int) []int {
	Outer: for i, firstNumber := range numbers {
		contiguousSet := []int {firstNumber}

		for j, additionalNumber := range numbers {
			if j<=i {
				continue
			}

			currentSum := sum(contiguousSet)
			if currentSum + additionalNumber == targetSum {
				contiguousSet = append(contiguousSet, additionalNumber)
				return contiguousSet
			} else if currentSum + additionalNumber < targetSum {
				contiguousSet = append(contiguousSet, additionalNumber)
			} else {
				continue Outer
			}
		}

	}

	return []int{}
}



func sum(numbers []int) (ret int) {
	for _, number := range numbers {
		ret += number
	}
	return
}

func findFaultyNumber(numbers []int) int {
	for i, number := range numbers {
		if i < preamble {
			continue
		}

		subset := numbers[i-preamble : i]
		validNumber := twoDifferentAddsTo(subset, number)

		if !validNumber {
			return number
		}

	}

	return 1
}

func twoDifferentAddsTo(numbers []int, targetSum int) bool {
	for i, number1 := range numbers {
		for j, number2 := range numbers {
			if i==j {
				continue
			}

			if number1 + number2 == targetSum {
				return true
			}

		}
	}

	return false
}

func max(numbers []int) int {
	ret := math.MinInt32
	for _, number := range numbers {
		if number > ret {
			ret = number
		}
	}
	return ret
}

func min(numbers []int) int {
	ret := math.MaxInt32
	for _, number := range numbers {
		if number < ret {
			ret = number
		}
	}
	return ret
}

func parseInput() (ret []int) {
	lines := utils.ReadLinesFromFile(fileName)
	for _, line := range lines {
		number, _ := strconv.Atoi(line)
		ret = append(ret, number)
	}
	return
}