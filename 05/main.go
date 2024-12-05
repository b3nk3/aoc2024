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

// sumOfValidMiddles returns the sum of the middle numbers
// in the sequences that are valid
func sumOfValidMiddles(rules map[int][]int, sequences [][]int) int {
	sum := 0
	for _, sequence := range sequences {
		if validateOrder(rules, sequence) {
			middle := (len(sequence) - 1) / 2
			sum += sequence[middle]

		}
	}
	return sum
}

func main() {
	data, err := os.ReadFile("input.txt")
	checkError(err)

	rules, orders := splitOdersAndRules(string(data))

	sum := sumOfValidMiddles(rules, orders)
	fmt.Println(sum)
}
