package main

import (
	"fmt"
	"../utils"
	"strings"
	"strconv"
	"log"
)

const fileName = "boot_code.txt"

func main() {

	instructions := parseInstructions()
	//fmt.Println(instructions)

	runUntilRepetition(instructions)

	findFaultyInstruction(instructions)

}

func findFaultyInstruction(instructions []instruction) {

	for i, instr := range instructions {
		var currentInstructionSet = []instruction{}
		if instr.operation == "jmp" {
			currentInstructionSet = changeOperationOnPosition(instructions, "nop", i)
		} else if instr.operation == "nop" {
			currentInstructionSet = changeOperationOnPosition(instructions, "jmp", i)
		} else {
			currentInstructionSet = changeOperationOnPosition(instructions, instr.operation, i)
		}
		
		ranToCompletion, globalAcc := tryRunToCompletion(currentInstructionSet)

		if ranToCompletion {
			fmt.Printf("Ran to completion. The global acc variable is %d\n", globalAcc)
			break
		}
	}

}

func changeOperationOnPosition(instructions []instruction, newOperation string, index int) []instruction {
	newInstructions := make([]instruction, len(instructions))

	copy(newInstructions, instructions)
	//fmt.Printf("Copied %d/%d elements. New instructions now contains %d elements\n", copiedElements, len(instructions), len(newInstructions))

	currentInstruction := instructions[index]
	newInstructions[index] = instruction{newOperation, currentInstruction.argument, index}
	return newInstructions
}

func tryRunToCompletion(instructions []instruction) (bool, int) {
	var visitedIndices = []int {}
	currentIndex := 0
	globalAcc := 0
	maxIndex := len(instructions) -1 

	visited := false
	for !visited {
		visitedIndices = utils.AppendIntIfNotPresent(visitedIndices, currentIndex)
		currentInstruction := instructions[currentIndex]
		nextIndex := currentInstruction.execute(&globalAcc)

		_, visited = utils.FindInt(visitedIndices, nextIndex)

		if nextIndex == maxIndex + 1 {
			return true, globalAcc
		}

		currentIndex = nextIndex
	}

	return false, globalAcc

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

		_, visited = utils.FindInt(visitedIndices, nextIndex)

		currentIndex = nextIndex
	}

	fmt.Printf("The global acc variable is %d\n", globalAcc)

}


func parseInstructions() (ret []instruction) {
	lines := utils.ReadLinesFromFile(fileName)
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

