package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func GetDirectionAndDistance(a, b int) (string, int) {
	if a == b {
		return "equal", 0
	}

	if a > b {
		return "decrease", a - b
	}

	return "increase", b - a
}

func IsSafeDistance(a, b int) bool {
	return math.Abs(float64(a-b)) < 4 && math.Abs(float64(a-b)) > 0
}

func CompareNumbers(a, b int) string {
	if a == b {
		return "equal"
	}

	if a > b {
		return "decrease"
	}

	return "increase"
}

func ReportAnalyzer(levels []int) bool {
	if levels == nil || len(levels) < 2 {
		return false
	}

	initialDistance := IsSafeDistance(levels[0], levels[1])
	if !initialDistance {
		return false
	}

	initialDirection := CompareNumbers(levels[0], levels[1])
	if initialDirection == "equal" {
		return false
	}

	prevLevel := levels[0]

	// start from the second element
	for _, level := range levels[1:] {
		// The levels are either all increasing or all decreasing.
		if initialDirection != CompareNumbers(prevLevel, level) {
			return false
		}

		// Any two adjacent levels differ by at least one and at most three.
		if !IsSafeDistance(level, prevLevel) {
			return false
		}

		prevLevel = level
	}
	return true
}

func Remove(a []int, index int) []int {
	return append(a[:index], a[index+1:]...)
}

func ProblemDampener(levels []int) bool {
	if levels == nil || len(levels) < 2 {
		return false
	}

	// First check if it's already safe without removing any numbers
	if ReportAnalyzer(levels) {
		return true
	}

	// Try removing each number one at a time and check if it becomes safe
	for i := range levels {
		// Make a copy of the slice without the current number
		testLevels := make([]int, len(levels)-1)
		copy(testLevels, levels[:i])
		copy(testLevels[i:], levels[i+1:])

		// testLevels := Remove(levels, i)
		fmt.Println(testLevels, levels)
		if ReportAnalyzer(testLevels) {
			return true
		}
	}

	return false
}

func main() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	// Split into lines
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	safeReports := 0
	for _, line := range lines {
		if line == "" {
			continue
		}

		// Split line into numbers
		numStrings := strings.Fields(line)
		numbers := make([]int, len(numStrings))

		for i, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				log.Fatal("Error converting string to int:", err)
			}
			numbers[i] = num
		}

		if ProblemDampener(numbers) {
			safeReports++
		}
	}

	fmt.Println("Safe reports:", safeReports)
}
