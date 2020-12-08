package main

import (
	"fmt"
	"../utils"
	"strings"
	"strconv"
	"log"
)

func main() {

	instructions := parseInstructions()
	//fmt.Println(instructions)

	runUntilRepetition(instructions)

}

func runUntilRepetition(instructions []instruction) {
	var visitedIndices = []int {}
	currentIndex := 0
	globalAcc := 0

	visited := false
	for !visited {
		visitedIndices = utils.AppendIntIfNotPresent(visitedIndices, currentIndex)
		currentInstruction := instructions[currentIndex]
		nextIndex := currentInstruction.execute(&globalAcc)
		currentIndex = nextIndex

		_, visited = utils.FindInt(visitedIndices, nextIndex)
	}

	fmt.Printf("The global acc variable is %d\n", globalAcc)

}


func parseInstructions() (ret []instruction) {
	lines := utils.ReadLinesFromFile("boot_code.txt")
	for i, line :=  range lines {
		instr := strings.Split(line, " ")
		operation := instr[0]
		argument, _ := strconv.Atoi(instr[1])
		ret = append(ret, instruction{operation, argument, i})
	}
	return
}

type instruction struct {
	operation string
	argument, index int
}

// Returns index of next code to be executed
func (i instruction) execute(globalAcc *int) int {
	if i.operation == "jmp" {
		return i.index + i.argument
	} else if i.operation == "acc" {
		*globalAcc += i.argument
		return i.index + 1
	} else if i.operation == "nop" {
		return i.index + 1
	} else {
		log.Fatal("Op code not found")
	}
	return -1
}

