package main

import (
	"fmt"
	"../utils"
	"errors"
	"sort"
)


func main() {

	lines := utils.ReadLinesFromFile("boarding_passes.txt")

	seats := findAllSeats(lines)

	maxID := maxID(seats)

	fmt.Printf("Max id is %d and min id is %d\n", maxID, minID(seats))

	sortedIds := sortedIds(seats)

	missingID := findMissingID(sortedIds)
	fmt.Printf("Missing id is %d\n", missingID)

}

type section  struct {
	min, max int
}

type seat struct {
	row, col int
}

func findMissingID(ids []int) int {

	for i:=1 ; i<len(ids) ; i++ {
		previous := ids[i-1]
		current := ids[i]
		if current-previous > 1 {
			return current -1
		}
	}

	panic(errors.New("Did not find missing id"))
}

func sortedIds(seats []seat) (ret []int) {
	for _, seat := range seats {
		ret = append(ret, seat.id())
	}

	sort.Ints(ret)
	fmt.Print(ret)
	return
}

func findAllSeats(boardingPasses []string) (ret []seat) {
	for _, boardingPass := range boardingPasses {
		ret = append(ret, findSeat(boardingPass))
	}
	return
}

func maxID(seats []seat) (ret int) {
	for _, seat := range seats {
		if seat.id() > ret {
			ret = seat.id()
		}
	}
	return
}

func minID(seats []seat) (ret int) {
	ret = 3000
	for _, seat := range seats {
		if seat.id() < ret {
			ret = seat.id()
		}
	}
	return
}

func findSeat(boardingPass string) seat {
	rowSection := section {0, 127}
	for i := 0; i<7; i++ {
		rowSection = rowSection.reduce(string(boardingPass[i]))
	}

	if rowSection.min != rowSection.max {
		panic(errors.New("Differing row section values after reduction"))
	}

	colSection := section {0, 7}
	for i := 7; i<10; i++ {
		colSection = colSection.reduce(string(boardingPass[i]))
	}

	if colSection.min != colSection.max {
		panic(errors.New("Differing col section values after reduction"))
	}

	//fmt.Printf("Found seat on row %d and col %d\n", rowSection.min, colSection.min)
	return seat{rowSection.min, colSection.min}
}

func (s seat) id() int {
	return s.row * 8 + s.col
}

func (s section) reduce(reductionLetter string) section {
	currentRange := (s.max - s.min) + 1
	if reductionLetter == "F" || reductionLetter == "L" {
		return section{s.min, s.min + (currentRange/2 - 1)}
	} else if reductionLetter == "B" || reductionLetter == "R" {
		return section{s.max - (currentRange/2 - 1), s.max}
	}
	panic(errors.New("Unknow direction"))
}