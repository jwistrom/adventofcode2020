package main

import (
	"fmt"
	"../utils"
)

func main() {
	deltaCoords := []coordinate {coordinate{1,1}, coordinate{1,3}, coordinate{1,5}, coordinate{1,7}, coordinate{2, 1}}
	
	for _, deltaCoord := range deltaCoords {
		treeCount := findEncounteredTrees(deltaCoord)
		fmt.Printf("Hit %d trees\n", treeCount)
	}

}

func findEncounteredTrees(deltaCoord coordinate) int {
	//fmt.Printf("\n\nDelta coord %d, %d\n", deltaCoord.row, deltaCoord.col)
	coord := coordinate{1, 1}

	forrestLines := utils.ReadLinesFromFile("map.txt")

	treeCount := 0
	lineIndex := 0

	for lineIndex < len(forrestLines) {
		forrestLine := forrestLine{forrestLines[lineIndex]}
		landedOnTree := forrestLine.hasTreeOn(coord)

		if (landedOnTree) {
			treeCount++
		}

		coord = coord.move(deltaCoord)
		lineIndex += deltaCoord.row
	}

	return treeCount
}


// one-index based
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
	lineLength := len(f.line)
	mod := coordinate.col % lineLength

	var positionInLine int

	if mod == 0 {
		positionInLine = lineLength
	} else {
		positionInLine = mod
	}

	return string(f.line[positionInLine-1]) == "#"
}

