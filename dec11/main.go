package main

import (
	"fmt"
	"../utils"
)

 const fileName = "map.txt"

 func main() {

	seats := utils.ReadLinesFromFile(fileName)

	
	part1(seats)
	part2(seats)

 }

 func part1(seats []string) {
	 changes := 1
	 iterations := 0
	 for changes > 0 {
		seats, changes = iterateSeatMap(seats)
		iterations++
		fmt.Printf("Iteration: %d. Changes was %d\n", iterations, changes)
	 }
	 fmt.Printf("The process ended with %d occupied seats\n", countOccupiedSeats(seats))
 }

 func part2(seats []string) {
	changes := 1
	iterations := 0
	for changes > 0 {
	   seats, changes = iterateSeatMapPart2(seats)
	   iterations++
	   fmt.Printf("Iteration: %d. Changes was %d\n", iterations, changes)
	}
	fmt.Printf("The process (part2) ended with %d occupied seats\n", countOccupiedSeats(seats))
 }

 func iterateSeatMap(seats []string) ([]string, int) {
	var updatedSeatMap []string
	var changes int
	for rowNum, seatRow := range seats {
		updatedSeatRow := ""
		for colNum, col := range seatRow {
		   seat := string(col)
	   
		   if seat == "L" && occupiedAdjacentSeats(seats, rowNum, colNum) == 0 {
			   updatedSeatRow = updatedSeatRow + "#"
			   changes++
		   } else if seat == "#" && occupiedAdjacentSeats(seats, rowNum, colNum) >= 4 {
			   updatedSeatRow = updatedSeatRow + "L"
			   changes++
		   } else {
			   updatedSeatRow = updatedSeatRow + seat
		   }
		}
		updatedSeatMap = append(updatedSeatMap, updatedSeatRow)
	}

	return updatedSeatMap, changes
}

 func iterateSeatMapPart2(seats []string) ([]string, int) {
	 var updatedSeatMap []string
	 var changes int
	 for rowNum, seatRow := range seats {
		 updatedSeatRow := ""
		 for colNum, col := range seatRow {
			seat := string(col)
		
			if seat == "L" && occupiedSeatsInSight(seats, rowNum, colNum) == 0 {
				updatedSeatRow = updatedSeatRow + "#"
				changes++
			} else if seat == "#" && occupiedSeatsInSight(seats, rowNum, colNum) >= 5 {
				updatedSeatRow = updatedSeatRow + "L"
				changes++
			} else {
				updatedSeatRow = updatedSeatRow + seat
			}
		 }
		 updatedSeatMap = append(updatedSeatMap, updatedSeatRow)
	 }

	 return updatedSeatMap, changes
 }


 func nextSeatOnLeftUpperDiagonal(row int, col int) (int, int) {
	 return row-1, col-1
 }

 func nextSeatOnRightUpperDiagonal(row int, col int) (int, int) {
	return row-1, col+1
}

func nextSeatOnLeftLowerDiagonal(row int, col int) (int, int) {
	return row+1, col-1
}

func nextSeatOnRightLowerDiagonal(row int, col int) (int, int) {
	return row+1, col+1
}

func nextSeatToRight(row int, col int) (int, int) {
	return row, col+1
}

func nextSeatToLeft(row int, col int) (int, int) {
	return row, col-1
}

func nextSeatUp(row int, col int) (int, int) {
	return row - 1, col
}

func nextSeatDown(row int, col int) (int, int) {
	return row + 1, col
}

 func occupiedSeatsInSight(seats []string, rowIdx int, colIdx int) (ret int) {
	rowLength := len(seats[0])

	functions := []func(int, int) (int, int) {nextSeatOnLeftUpperDiagonal, nextSeatOnRightUpperDiagonal, nextSeatOnLeftLowerDiagonal, nextSeatOnRightLowerDiagonal,
		nextSeatDown, nextSeatUp, nextSeatToLeft, nextSeatToRight}

	for _, movingFunction := range functions {
		currentRow, currentCol := movingFunction(rowIdx, colIdx)
		for currentCol >= 0 && currentCol < rowLength && currentRow >=0 && currentRow < len(seats) {
			currentSeat := seatOnPosition(seats, currentRow, currentCol)
			if currentSeat == "#" {
				ret++
				break
			} else if currentSeat == "L" {
				break
			}
			currentRow, currentCol = movingFunction(currentRow, currentCol)
		}
	}
	return

 }

 func occupiedAdjacentSeats(seats []string, row int, col int) (ret int) {
	 rowLength := len(seats[0])
	 for r:=row-1 ; r<=row+1 ; r++ {
		 if r < 0 || r>= len(seats) {
			 continue
		 }
		 for c:=col-1 ; c<=col+1 ; c++ {
			 if c < 0 || c>= rowLength || (c==col && r==row) {
				 continue
			 }

			 seat := string(seats[r][c])
			 if seat == "#" {
				 ret++
			 }

		 }
	 }

	 return
 }

 func countOccupiedSeats(seats []string) (ret int) {
	 for _, row := range seats {
		 for _, col := range row {
			seat :=  string(col)
			if seat == "#" {
				ret++
			}
		 }
	 }
	 return
 }

 func seatOnPosition(seats []string, row int, col int) string {
	return string(seats[row][col])
 }

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}