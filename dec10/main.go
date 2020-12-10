package main

import (
	"fmt"
	"../utils"
	"strconv"
	"sort"
	"math"
	"errors"
)

const fileName = "test_joltage.txt"

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



	
	//reverseOrderSort := sort.Reverse(sort.IntSlice(adapters))
	//sort.Sort(reverseOrderSort)
	combinations := 1
	fmt.Println(adapters)
	findAllCombinations4(adapters, &combinations, 1, 1, len(adapters)-2, -1)

	fmt.Printf("Combinations: %d\n", combinations)
}

func findAllCombinations4(allAdapters []int, combinations *int, level int, minIndex int, maxIndex int, nextToStart int) {
	
	for i := 1 ; i < len(allAdapters)-1 ; i++ {
		current := allAdapters[i]
		if nextToStart != -1 && nextToStart != current {
			 continue
		} 
		

		numberCanBeRemoved := canBeRemoved(allAdapters, current)
		if numberCanBeRemoved {
			*combinations = *combinations + 1
			
			findAllCombinations4(removeFromSlice(allAdapters, current), combinations, level+1, i, i, allAdapters[i+1])
		} else if level > 1 {
			return
		}

	}

}

func canBeRemoved(slice []int, value int) bool {
	for i, number := range slice {
		if number == value {
			return slice[i-1] - slice[i+1] <= 3
		}
	}
	panic(errors.New("Could not find number"))
}

func removeFromSlice(slice []int, value int) (ret []int){
	for _, number := range slice {
		if number != value {
			ret = append(ret, number)
		}
	}
	return
}

func findAllCombinations3(allAdapters []int) {
	sort.Ints(allAdapters)
	fmt.Println(allAdapters)
	
	var waysToGetToIndex []int
	for i, targetAdapt := range allAdapters {
		if i == 0 {
			waysToGetToIndex = append(waysToGetToIndex, 1)
			continue
		}

		var waysToTarget int
		for j, potentialSourceAdapt := range allAdapters {
			if j >= i {
				break
			}

			isValidSource := targetAdapt - potentialSourceAdapt <= 3
			if isValidSource {
				waysToTarget++
			} else {
				continue
			}
		}
		waysToGetToIndex = append(waysToGetToIndex, waysToTarget)
	}

	fmt.Println(waysToGetToIndex)
	fmt.Printf("Number of combinations are %d\n", multiply(waysToGetToIndex))
	
}

func findAllCombinations2(allAdapters []int) {
	sort.Ints(allAdapters)
	fmt.Println(allAdapters)

	reverseOrderSort := sort.Reverse(sort.IntSlice(allAdapters))
	sort.Sort(reverseOrderSort)
	
	var waysToGetToIndex []int
	for i, targetAdapt := range allAdapters {
		if i == len(allAdapters) - 1 {
			break
		}

		var waysToTarget int
		for j, potentialSourceAdapt := range allAdapters {
			if j <= i {
				continue
			}

			isValidSource := targetAdapt - potentialSourceAdapt <= 3
			if isValidSource {
				waysToTarget++
			} else {
				break
			}
		}
		waysToGetToIndex = append(waysToGetToIndex, waysToTarget)
	}

	fmt.Println(waysToGetToIndex)
	fmt.Printf("Number of combinations are %d\n", multiply(waysToGetToIndex))
	
}

func multiply(numbers []int) int {
	ret := 1
	for _, number := range numbers {
		ret = ret * number
	}
	return ret
}

func findAllCombinations(allAdapters []int) {
	if !sort.IntsAreSorted(allAdapters) {
		panic(errors.New("Adapters must be sorted"))
	}

	combinationCount := 0
	for i, adapt := range allAdapters {
		if adapt > 3 {
			break
		}

		currentAdapters := []int {adapt}
		findNextAdapter(allAdapters, currentAdapters, i+1, &combinationCount)
	}


	fmt.Printf("Number of combinations are %d\n", combinationCount)
}

func findNextAdapter(allAdapters []int, currentAdapters []int, startSearchAt int, combinationCount *int) {

	lastAdapter := currentAdapters[len(currentAdapters) - 1]

	for i, adapt := range allAdapters {
		if i < startSearchAt {
			continue
		}

		isValidNext := adapt - lastAdapter <= 3

		if !isValidNext {
			break
		}

		if startSearchAt == len(allAdapters) - 1 {
			*combinationCount = *combinationCount + 1
			//fmt.Printf("Found combination %v\n", currentAdapters)
			fmt.Printf("Found combinations %d\n", *combinationCount)
		}

		newCurrentAdapters := make([]int, len(currentAdapters))
		copy(newCurrentAdapters, currentAdapters)
		newCurrentAdapters = append(newCurrentAdapters, adapt)

		findNextAdapter(allAdapters, newCurrentAdapters, i+1, combinationCount)
	}
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