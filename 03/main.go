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

func main() {
	// Read input file
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}
	muls := FindMulsWithRegexp(string(data))
	fmt.Println(MultiPlySliceItems(muls))
}
