package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func bytesToStrings(data [][]byte) []string {
	result := make([]string, len(data))
	for i, b := range data {
		result[i] = string(b)
	}
	return result
}

func FindMulsWithRegexp(s string) []string {
	var re = regexp.MustCompile(`mul\(\d{1,3}\,\d{1,3}\)`)
	return bytesToStrings(re.FindAll([]byte(s), -1))
}

func MultiPlySliceItems(s []string) int {
	result := 0
	for _, v := range s {
		var a, b int
		fmt.Sscanf(v, "mul(%d,%d)", &a, &b)
		result += a * b
	}

	return result
}

func FilterOutBlocks(s string) string {
	// used `don't\(\).*?do\(\)` intitally, but it was not working as expected
	// so I used `don't\(\).*?do\(\)` and it worked because
	//
	// Without (?s), the . matches any character except newline
	// With (?s), the . matches absolutely any character including newlines
	// The *? keeps it non-greedy, so it matches the minimum needed between your delimiters
	re := regexp.MustCompile(`(?s)don't\(\).*?do\(\)`)
	parts := re.ReplaceAllString(s, "")
	return parts
}

func main() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	filteredMuls := FilterOutBlocks(string(data))
	muls := FindMulsWithRegexp(filteredMuls)
	fmt.Println(MultiPlySliceItems(muls))
}
