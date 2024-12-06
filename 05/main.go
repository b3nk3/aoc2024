package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		log.Fatal("Error: ", err)
		os.Exit(1)
	}
}

func splitOdersAndRules(input string) (map[int][]int, [][]int) {
	chunks := strings.Split(strings.TrimSpace(input), "\n\n")
	rules, orders := chunks[0], chunks[1]
	rulesMap := map[int][]int{}
	ordersMatrix := [][]int{}

	for _, rule := range strings.Split(rules, "\n") {
		rule := strings.Split(rule, "|")

		ruleKey, err := strconv.Atoi(rule[0])
		checkError(err)
		ruleVal, err := strconv.Atoi(rule[1])
		checkError(err)

		rulesMap[ruleKey] = append(rulesMap[ruleKey], ruleVal)
	}

	for _, order := range strings.Split(orders, "\n") {
		order := strings.Split(order, ",")
		orderInt := []int{}
		for _, o := range order {
			oInt, err := strconv.Atoi(o)
			checkError(err)
			orderInt = append(orderInt, oInt)
		}
		ordersMatrix = append(ordersMatrix, orderInt)
	}

	return rulesMap, ordersMatrix
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// validateOrder validates the order of the sequence
// based on the rules
// returns an array of valid sequence indices
func validateOrder(rules map[int][]int, sequence []int) bool {
	for i, num := range sequence {
		if i == 0 {
			continue
		}
		mustComeBefore := rules[num]
		// num must come before all the numbers in mustComeBefore
		// to determine we'll check if any of the previous numbers in the sequence
		// are in mustComeBefore
		// if so, return false
		for _, subNum := range sequence[:i] {
			if contains(mustComeBefore, subNum) {
				return false
			}
		}

	}

	return true
}

func getMiddleNumber(sequence []int) int {
	if len(sequence)%2 == 0 {
		return 0
	}
	middle := (len(sequence) - 1) / 2
	return sequence[middle]
}

// sumOfValidMiddlesAndInvalids returns the sum of the middle numbers
// in the sequences that are valid
func sumOfValidMiddlesAndInvalids(rules map[int][]int, sequences [][]int) (int, [][]int) {
	sum := 0
	invalidSequences := [][]int{}
	for _, sequence := range sequences {
		if validateOrder(rules, sequence) {
			sum += getMiddleNumber(sequence)
		} else {
			invalidSequences = append(invalidSequences, sequence)
		}
	}
	return sum, invalidSequences
}

func countInDegrees(rules map[int][]int) map[int]int {
	inDegrees := map[int]int{}

	for ruleKey := range rules {
		inDegrees[ruleKey] = 0 // Initialize source nodes too
	}

	for _, destinations := range rules {
		for _, dest := range destinations {
			inDegrees[dest]++
		}
	}
	return inDegrees
}

// fixOrder fixes the order of the sequence
func fixOrder(rules map[int][]int, sequence []int) []int {
	// Create adjacency list representation of the graph
	graph := make(map[int][]int)
	for _, num := range sequence {
		graph[num] = []int{}
	}

	// Build edges based on rules that apply to numbers in our sequence
	for _, num := range sequence {
		for _, mustComeBefore := range rules[num] {
			if contains(sequence, mustComeBefore) {
				graph[num] = append(graph[num], mustComeBefore)
			}
		}
	}

	// Create in-degree map for Kahn's algorithm
	inDegree := make(map[int]int)
	for _, num := range sequence {
		inDegree[num] = 0
	}
	for _, edges := range graph {
		for _, dest := range edges {
			inDegree[dest]++
		}
	}

	// Find nodes with no incoming edges
	var queue []int
	for _, num := range sequence {
		if inDegree[num] == 0 {
			queue = append(queue, num)
		}
	}

	// Process nodes in topological order
	var result []int
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		result = append(result, current)

		for _, neighbor := range graph[current] {
			inDegree[neighbor]--
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	return result
}

// sumOfFixedMiddles returns the sum of the middle numbers
// in the sequences that are fixed
func sumOfFixedMiddles(rules map[int][]int, invalidSequences [][]int) int {
	sum := 0
	for _, sequence := range invalidSequences {
		if len(sequence)%2 == 0 {
			continue // Skip even-length sequences as they don't have a middle
		}
		fixedSequence := fixOrder(rules, sequence)
		sum += getMiddleNumber(fixedSequence)
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	rules, orders := splitOdersAndRules(string(data))
	sum, invalidSequences := sumOfValidMiddlesAndInvalids(rules, orders)
	fmt.Println("Sum of valid middle numbers: ", sum)

	fixedSum := sumOfFixedMiddles(rules, invalidSequences)
	fmt.Println("Sum of fixed middle numbers: ", fixedSum)
}
