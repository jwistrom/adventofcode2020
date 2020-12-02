package main

import (
	"fmt"
	"errors"
	"strconv"
	"os"
	"bufio"
	"log"
)

func main() {
	
	expenses := readExpenses()

	a, b, err := findPairsSummingTo(expenses, 2020)
	check(err)

	fmt.Printf("The numbers %d and %d add to 2020 and the multiplication of the two yields %d\n", a, b, a*b)

	c, d, e, err := findTripletsSummingTo(expenses, 2020)
	check(err)

	fmt.Printf("The numbers %d, %d and %d add to 2020 and the multiplication of the two yields %d\n", c, d, e, c*d*e)


}

func findPairsSummingTo(expenses []int, sum int) (int, int, error) {
	for i, expense := range expenses {
		for j, reference := range expenses {
			if (j == i) {
				continue
			} else if (expense + reference == sum) {
				return expense, reference, nil
			}
		}
	}
	log.Fatal("No pairs adding to " + string(sum))
	return 0, 0, errors.New("No pairs adding to " + string(sum))
}

func findTripletsSummingTo(expenses []int, sum int) (int, int, int, error) {
	for i, expense1 := range expenses {
		for j, expense2 :=  range expenses {
			if (expense1 + expense2 >= sum || i==j) {
				continue
			}
			for k, expense3 := range expenses {
				if (k == i || k == j) {
					continue
				} else if (expense1 + expense2 + expense3 == sum) {
					return expense1, expense2, expense3, nil
				}
			}
		}
	}
	return 0, 0, 0, errors.New("No triplet found adding to "  + string(sum))
}

func readExpenses() []int {
	file, err := os.Open("expenses.txt")
	check(err)

	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())
		check(err)
		lines = append(lines, line)
	}

	check(scanner.Err())

	return lines
}


func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

