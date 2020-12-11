package main

import (
	"fmt"
	"../utils"
	"strconv"
	"sort"
	"math"
	"errors"
)

const fileName = "output_joltage.txt"

func main() {
	adapters := parseInput()

	sort.Ints(adapters)
	fmt.Println(adapters)

	deviceJoltage := max(adapters) + 3
	adaptersWithDevice := make([]int, len(adapters))
	copy(adaptersWithDevice, adaptersWithDevice)

	adaptersWithDevice = append(adapters, deviceJoltage)

	distribution := findJoltageDifferenceDistribution(adaptersWithDevice)
	fmt.Println(distribution)
	fmt.Printf("The answer to part 1 is %d\n", distribution[1] * distribution[3])

	allAdapters := make([]int, len(adaptersWithDevice))
	copy(allAdapters, adaptersWithDevice)
	allAdapters = append(allAdapters, 0)
	sort.Ints(allAdapters)
	
	combinations := findCombinationsCount(allAdapters)
	fmt.Printf("The answer to part 2 is %d\n", combinations)
}

func findCombinationsCount(allAdapters []int) int {
	combinationCount := map[int]int {0:1}
	combinationCount[0] = 1

	for _, adapt := range allAdapters {
		combinationCount[adapt] += combinationCount[adapt-1]
		combinationCount[adapt] += combinationCount[adapt-2]
		combinationCount[adapt] += combinationCount[adapt-3]
	}

	return getLast(combinationCount)
}

func getLast(m map[int]int) (ret int){

	var keys []int
	for k := range m {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		ret = m[k]
	}
	return
}

func findJoltageDifferenceDistribution(adapters []int) map[int]int {

	sort.Ints(adapters)
	ret := make(map[int]int)
	
	currentInput := 0
	for _, currentAdapter := range adapters {
		difference := currentAdapter - currentInput
		if difference > 3 {
			panic(errors.New("The joltage difference exceeds 3"))
		}

		ret[difference] = ret[difference] + 1
		currentInput = currentAdapter

	}

	return ret

}

func parseInput() (ret []int) {
	lines := utils.ReadLinesFromFile(fileName)
	for _, line := range lines {
		joltage, _ := strconv.Atoi(line)
		ret = append(ret, joltage)
	}
	return
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