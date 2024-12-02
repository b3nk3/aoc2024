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

func main() {
	data, err := os.ReadFile("input.tsv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	// Split into lines
	lines := strings.Split(string(data), "\n")

	safeReports := 0
	for _, line := range lines {
		if err != nil {
			log.Fatal(err)
		}
		// Split line into numbers
		numStrings := strings.Fields(line)

		var numbers []int
		for _, numStr := range numStrings {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Error converting string to int:", err)
				os.Exit(1)
			}
			numbers = append(numbers, num)
		}
		if ReportAnalyzer(numbers) {
			safeReports++
		}
	}
	fmt.Println(safeReports)
}
