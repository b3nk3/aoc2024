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
		inDegrees[ruleKey] = 0
		for _, num := range rules[ruleKey] {
			inDegrees[num] = 0
		}
	}
	for ruleKey := range rules {
		for _, num := range rules[ruleKey] {
			inDegrees[num] += 1
		}
	}

	return inDegrees
}

// fixOrder fixes the order of the sequence
func fixOrder(rules map[int][]int, sequence []int) []int {
	inDegrees := countInDegrees(rules)

	// make a copy of the sequence
	// to prevent mutation
	// I didn't realise this, and without my tests I wouldn't have known
	remainingInSequence := make([]int, len(sequence))
	copy(remainingInSequence, sequence)

	fixedSequence := []int{}

	// for len(remainingInSequence) > 0 {
	for i := len(remainingInSequence); i > 0; i-- {
		zeroDegrees := []int{}

		for degree := range inDegrees {
			if inDegrees[degree] == 0 {
				zeroDegrees = append(zeroDegrees, degree)
			}
		}

		fixedSequence = append(fixedSequence, zeroDegrees...)
		// decrease the in-degree of the numbers that depend on the number
		for _, degree := range zeroDegrees {
			for _, num := range rules[degree] {
				inDegrees[num] -= 1
			}
			// delete the number from the in-degrees
			delete(inDegrees, degree)
		}

		// remove the fixed numbers from the sequence
		for i, num := range remainingInSequence {
			if contains(fixedSequence, num) {
				remainingInSequence = append(remainingInSequence[:i], remainingInSequence[i+1:]...)
			}
		}
	}
	return fixedSequence
}

// sumOfFixedMiddles returns the sum of the middle numbers
// in the sequences that are fixed
func sumOfFixedMiddles(rules map[int][]int, invalidSequences [][]int) int {
	sum := 0
	fixedSequence := []int{}
	for _, sequence := range invalidSequences {
		fmt.Println("Sequence: ", sequence)
		fixedSequence = fixOrder(rules, sequence)
		fmt.Println("Fixed sequence: ", fixedSequence)
		sum += getMiddleNumber(fixedSequence)
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	rules, orders := splitOdersAndRules(string(data))

	sum, invalidSequences := sumOfValidMiddlesAndInvalids(rules, orders)

	fixedSum := sumOfFixedMiddles(rules, invalidSequences)

	fmt.Println("Sum of valid middle numbers: ", sum)
	fmt.Println("Sum of fixed middle numbers: ", fixedSum)
}
