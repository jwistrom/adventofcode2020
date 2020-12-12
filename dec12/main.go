package main

import (
	"fmt"
	"../utils"
	"strconv"
	"errors"
)


const fileName = "test_input.txt"

func main() {

	instructions := parseInstructions()
	part1(instructions)

}

func part1(instructions []instruction) {
	currentPosition := position{0, 0, "E"}

	for _, instruction := range instructions {
		currentPosition = currentPosition.move(instruction)
	}

	manhattanDistance := utils.Abs(currentPosition.ns) + utils.Abs(currentPosition.ew)
	fmt.Printf("Answer to part 1: %d\n", manhattanDistance)
}

func parseInstructions() (ret []instruction) {
	lines := utils.ReadLinesFromFile(fileName)

	for _, line := range lines {
		dir := direction(string(line[0]))
		amount, _ := strconv.Atoi(line[1:])
		ret = append(ret, instruction{dir, amount})
	}
	return
}

type direction string

type instruction struct {
	direction direction
	amount int
}

type position struct {
	ns, ew int
	heading direction
}

func (p position) move(instruct instruction) position {
	if instruct.direction == "N" {
		return position{p.ns + instruct.amount, p.ew, p.heading}
	} else if instruct.direction == "S" {
		return position{p.ns - instruct.amount, p.ew, p.heading}
	} else if instruct.direction == "E" {
		return position{p.ns, p.ew + instruct.amount, p.heading}
	} else if instruct.direction == "W" {
		return position{p.ns, p.ew - instruct.amount, p.heading}
	} else if instruct.direction == "F" {
		proxyInstruction := instruction{p.heading, instruct.amount}
		return p.move(proxyInstruction)
	} else if instruct.direction == "L" || instruct.direction == "R" {
		newHeading := p.heading.rotate(instruct.amount, instruct.direction)
		return position{p.ns, p.ew, newHeading}
	}
	panic(errors.New("Unknow instruction"))
}

func (d direction) rotate(degrees int, rotation direction) direction {
	var newHeading int
	if rotation == "L" {
		newHeading  = utils.Mod(d.intValue() - degrees, 360)
	} else if rotation == "R" {
		newHeading  = utils.Mod(d.intValue() + degrees, 360)
	}

	if newHeading < 0 {
		fmt.Printf("Prev heading was %s. will rotate %d degrees in direct %s. Resulted in %d\n", d, degrees, rotation, newHeading)
	}

	return directionFromInt(newHeading)
}

func (d direction) intValue() int {
	if d == "N" {
		return 0
	} else if d == "E" {
		return 90
	} else if d == "S" {
		return 180
	} else if d == "W" {
		return 270
	}

	panic(errors.New("Invalid direction"))
}

func directionFromInt(i int) direction {
	if i == 0 {
		return "N"
	} else if i == 90{
		return "E"
	} else if i == 180 {
		return "S"
	} else if i == 270 {
		return "W"
	}

	panic(errors.New("Invalid degrees " + fmt.Sprint(i)))
}