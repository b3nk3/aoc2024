package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
}

func checkMatrix(matrix [][]string, directions [][]int, row, target []string, start_x, start_y int) int {
	numberOfOcurrences := 0
	for _, direction := range directions {
		x, y := start_x, start_y
		dx, dy := direction[0], direction[1]
		targetWord := target[1:]
		for z, letter := range targetWord {
			x += dx
			y += dy
			if x < 0 || x >= len(matrix) || y < 0 || y >= len(row) {
				break
			}
			if matrix[x][y] != letter {
				break
			}
			if z == len(targetWord)-1 {
				numberOfOcurrences++
			}
		}
	}
	return numberOfOcurrences
}

// findXmas returns the number of times the word "XMAS" appears in the matrix
func findXmas(matrix [][]string) (int, error) {
	if matrix == nil {
		return 0, errors.New("missing matrix")
	}
	target1 := []string{"X", "M", "A", "S"}
	target2 := []string{"S", "A", "M", "X"}

	directions := [][]int{
		{-1, 1}, // diagonal up-right
		{0, 1},  // right
		{1, 1},  // diagonal down-right
		{1, 0},  // down
	}
	// iterate through the matrix
	// check if the word "XMAS" appears in the matrix
	// return the number of times the word "XMAS" appears in the matrix

	totalOccurrences := 0

	for i, row := range matrix {
		for j, char := range row {

			if char == target1[0] {
				start_x, start_y := i, j
				totalOccurrences += checkMatrix(matrix, directions, row, target1, start_x, start_y)
			}
			if char == target2[0] {
				start_x, start_y := i, j
				totalOccurrences += checkMatrix(matrix, directions, row, target2, start_x, start_y)
			}
		}
	}
	return totalOccurrences, nil
}

type Direction string

const (
	Forward  Direction = "forward"
	Backward Direction = "backward"
)

func isValidAxis(matrix [][]string, directions map[string][]int, direction Direction, row []string, start_x, start_y int) (string, string) {
	x, y, a, b := 0, 0, 0, 0
	switch direction {
	case Forward:
		x, y = start_x+directions["upRight"][0], start_y+directions["upRight"][1]
		a, b = start_x+directions["downLeft"][0], start_y+directions["downLeft"][1]
	case Backward:
		a, b = start_x+directions["upLeft"][0], start_y+directions["upLeft"][1]
		x, y = start_x+directions["downRight"][0], start_y+directions["downRight"][1]
	}

	if x < 0 || x >= len(matrix) || y < 0 || y >= len(row) {
		return "", ""
	}
	if a < 0 || a >= len(matrix) || b < 0 || b >= len(row) {
		return "", ""
	}
	topLetter := matrix[x][y]
	bottomLetter := matrix[a][b]

	return topLetter, bottomLetter
}

// findCrossMas returns the number of times the word "X-MAS" appears in the matrix
// X shaped MAS, that is
//
// M.S
// .A.
// M.S
func findCrossMas(matrix [][]string) (int, error) {
	if matrix == nil {
		return 0, errors.New("missing matrix")
	}

	directions := map[string][]int{
		"upRight":   {-1, 1},  // diagonal up-right
		"downRight": {1, 1},   // diagonal down-right
		"upLeft":    {-1, -1}, // diagonal up-left
		"downLeft":  {1, -1},  // diagonal down-left
	}

	totalXOccurrences := 0

	// iterate through the matrix
	for i, row := range matrix {

		if i == 0 || i == len(matrix)-1 {
			continue
		}
		for j, char := range row {
			if j == 0 || j == len(row)-1 {
				continue
			}
			if char == "A" {
				upRightLetter, downLeftLetter := isValidAxis(matrix, directions, "forward", row, i, j)

				if upRightLetter == downLeftLetter {
					continue
				}

				if (upRightLetter == "M" && downLeftLetter == "S") || (upRightLetter == "S" && downLeftLetter == "M") {
					downRightLetter, upLeftLetter := isValidAxis(matrix, directions, "backward", row, i, j)

					if downRightLetter == upLeftLetter {
						continue
					}

					if (downRightLetter == "M" && upLeftLetter == "S") || (downRightLetter == "S" && upLeftLetter == "M") {
						totalXOccurrences++
					}
				}

			}

		}
	}

	return totalXOccurrences, nil
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	wordMatrix := [][]string{}
	for _, line := range strings.Split(string(data), "\n") {
		if line != "" {
			// Split line into individual characters
			chars := strings.Split(line, "")
			wordMatrix = append(wordMatrix, chars)
		}
	}
	// task 1
	number, err := findXmas(wordMatrix)
	checkError(err)
	fmt.Println("Task 1 solution: ", number)

	// task 2
	number, err = findCrossMas(wordMatrix)
	checkError(err)
	fmt.Println("Task 2 solution: ", number)

}
