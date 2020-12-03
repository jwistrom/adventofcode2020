package main

import (
	"fmt"
	"../utils"
)

func main() {
	deltaCoord := coordinate{row: 1, col: 3}
	coord := coordinate{0, 0}

	forrestLines := utils.ReadLinesFromFile("map.txt")

	treeCount := 0
	for i, line := range forrestLines {
		forrestLine := forrestLine{line}
		landedOnTree := forrestLine.hasTreeOn(coord)

		fmt.Printf("On line %d, we are on coord row: %d and col: %d and hit tree: %t\n", i, coord.row, coord.col, landedOnTree)

		if (landedOnTree) {
			treeCount++
		}

		coord = coord.move(deltaCoord)
	}

	fmt.Printf("Hit %d trees\n", treeCount)

}


// zero-index based
type coordinate struct {
	row, col int
}

func (c coordinate) move(delta coordinate) coordinate {
	return coordinate{row: c.row + delta.row, col: c.col + delta.col}
}

type forrestLine struct {
	line string
}

func (f forrestLine) hasTreeOn(coordinate coordinate) bool {
	lineSize := len(f.line)
	soughtIndex := coordinate.col % lineSize
	return string(f.line[soughtIndex]) == "#"
}

