package main

import (
	"reflect"
	"testing"
)

var input = `
47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`
var smallInput = `
47|53
97|13
97|61
75|47

75,47,61,53,29
97,61,53,29,13
75,29,13
`

var ruleSet1 = map[int][]int{
	47: {53},
	75: {47},
	97: {13, 61},
}

var orderSet1 = [][]int{
	{75, 47, 61, 53, 29},
	{97, 61, 53, 29, 13},
	{75, 29, 13},
}

var ruleSet2 = map[int][]int{
	29: {13},
	47: {53, 13, 61, 29},
	53: {29, 13},
	61: {13, 53, 29},
	75: {29, 53, 47, 61, 13},
	97: {13, 61, 47, 29, 53, 75},
}

var orderSet2 = [][]int{
	{75, 47, 61, 53, 29},
	{97, 61, 53, 29, 13},
	{75, 29, 13},
	{75, 97, 47, 61, 53},
	{61, 13, 29},
	{97, 13, 75, 29, 47},
}

func TestSumOfValidOrders(t *testing.T) {
	t.Run("should return the sum of valid orders", func(t *testing.T) {
		rules, orders := splitOdersAndRules(input)

		got := sumOfValidMiddles(rules, orders)
		wanted := 143

		if got != wanted {
			t.Errorf("got %v, wanted %v", got, wanted)
		}
	})
}

func TestValidateOrder(t *testing.T) {

	t.Run("should return true for valid order", func(t *testing.T) {
		rules := ruleSet1
		order := orderSet1[0]

		got := validateOrder(rules, order)
		wanted := true

		if got != wanted {
			t.Errorf("got %v, wanted %v", got, wanted)
		}
	})
	t.Run("should return false for invalid order", func(t *testing.T) {
		rules := ruleSet2
		order := orderSet2[3]

		got := validateOrder(rules, order)
		wanted := false

		if got != wanted {
			t.Errorf("got %v, wanted %v", got, wanted)
		}
	})
	t.Run("should return false for invalid order", func(t *testing.T) {
		rules := ruleSet2
		order := orderSet2[5]

		got := validateOrder(rules, order)
		wanted := false

		if got != wanted {
			t.Errorf("got %v, wanted %v", got, wanted)
		}
	})
}

func TestSplitOrderAndRules(t *testing.T) {
	t.Run("should split orders and rules", func(t *testing.T) {
		rules, orders := splitOdersAndRules(smallInput)

		wantedRules := ruleSet1

		wantedOrder := orderSet1

		if !reflect.DeepEqual(rules, wantedRules) {
			t.Errorf("got %v, wanted %v", rules, wantedRules)
		}
		if !reflect.DeepEqual(orders, wantedOrder) {
			t.Errorf("got %v, wanted %v", orders, wantedOrder)
		}

	})
	t.Run("should split orders and rules", func(t *testing.T) {
		rules, orders := splitOdersAndRules(input)

		wantedRules := ruleSet2
		wantedOrder := orderSet2

		if !reflect.DeepEqual(rules, wantedRules) {
			t.Errorf("got %v, wanted %v", rules, wantedRules)
		}
		if !reflect.DeepEqual(orders, wantedOrder) {
			t.Errorf("got %v, wanted %v", orders, wantedOrder)
		}

	})
}
