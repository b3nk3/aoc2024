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

	totalOcurrences := 0

	for i, row := range matrix {
		for j, char := range row {

			if char == target1[0] {
				start_x, start_y := i, j
				totalOcurrences += checkMatrix(matrix, directions, row, target1, start_x, start_y)
			}
			if char == target2[0] {
				start_x, start_y := i, j
				totalOcurrences += checkMatrix(matrix, directions, row, target2, start_x, start_y)
			}
		}
	}
	return totalOcurrences, nil
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
	number, err := findXmas(wordMatrix)
	checkError(err)

	fmt.Println(number)

}
