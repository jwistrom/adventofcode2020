package main

import (
	"fmt"
	"../utils"
)

 const fileName = "map.txt"

 func main() {

	seats := utils.ReadLinesFromFile(fileName)

	
	part1(seats)

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