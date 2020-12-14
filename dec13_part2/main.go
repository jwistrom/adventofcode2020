package main

import (
	"fmt"
	"../utils"
	"strconv"
	"strings"
)

const fileName = "input.txt"

type busID string

func main() {
	busIDs := parseInput()
	part2(busIDs)
}

func part2(busIDs []busID) {

	var t, step = uint64(0), uint64(1)
	for idx, bus := range busIDs {
		if bus.isAnyBus() {
			continue
		}

		for (t+uint64(idx))%bus.asInt() != 0 {
			t += step
		}

		step *= bus.asInt()
	}
	
	fmt.Printf("Departure time %d\n", t)
}

func parseInput() (ret []busID) {
	lines := utils.ReadLinesFromFile(fileName)

	for _, line := range strings.Split(lines[1], ",") {
		ret = append(ret, busID(line))
	}
	return
}

func (id busID) isAnyBus() bool {
	return id == "x"
}

func (id busID) asInt() uint64 {
	ret, _ := strconv.ParseUint(string(id), 10, 64)
	return ret
}


