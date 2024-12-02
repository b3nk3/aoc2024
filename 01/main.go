package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func SortNumberArrays(numbers []int) []int {
	slices.Sort(numbers)
	return numbers
}

func sumRange(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func GetTheTotalDistance(left, right []int) int {
	sortedLeftList := SortNumberArrays(left)
	sortedRightList := SortNumberArrays(right)

	var distances []int

	for i := 0; i < len(sortedLeftList); i++ {
		diff := sortedLeftList[i] - sortedRightList[i]
		distances = append(distances, int(math.Abs(float64(diff))))
	}
	return sumRange(distances)
}

func GetSimilarityScore(left, right []int) int {
	frequency := make(map[int]int)

	for i := 0; i < len(right); i++ {
		element := right[i]
		frequency[element] = frequency[element] + 1
	}

	similarity := 0

	for i := 0; i < len(left); i++ {
		element := left[i]
		appearance := frequency[element]
		similarity += element * appearance
	}

	return similarity
}

func main() {
	input, err := os.ReadFile("input.tsv")
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	var (
		leftArray, rightArray []int
	)

	r := csv.NewReader(bytes.NewReader(input))

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		parts := strings.Fields(record[0])

		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal(err)
		}

		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}

		leftArray = append(leftArray, leftNum)
		rightArray = append(rightArray, rightNum)

	}

	totalDistance := GetTheTotalDistance(leftArray, rightArray)
	fmt.Println("totalDistance: ", totalDistance)

	similarityScore := GetSimilarityScore(leftArray, rightArray)
	fmt.Println("similarityScore: ", similarityScore)
}
