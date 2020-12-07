package main

import (
	"aod2020/pkg/config"
)

// Rule is one line in the input rule document.
type Rule struct {
	// The name of the bag type that needs to follow the rule.
	bag     string
	// A list of bags that needs to be inside the bag following this rule.
	requiredContents []struct {
		// The name of the bag that needs to be contained.
		bag string
		// The amount of bags of this type that needs to be included.
		count int
	}
}

func (r Rule) Contains(bag string) bool {
	for _, b := range r.requiredContents {
		if b.bag == bag {
			return true
		}
	}
	return false
}

func main() {
	part, in := config.ParseDefaultFlags()
	rules := readRules(in)
	if part == 1 {
		println("The number of possible bags are:", walkConnectedRules(rules, "shiny gold", make(map[string]bool)))
	}
	if part == 2 {
		println("The golden bag contains:", countBags(rules, "shiny gold"))
	}
}

func walkConnectedRules(rules map[string]Rule, node string, visited map[string]bool) int {
	for _, r := range rules {
		if r.Contains(node) {
			visited[r.bag] = true
			walkConnectedRules(rules, r.bag, visited)
		}
	}
	return len(visited)
}

func countBags(rules map[string]Rule, node string) int {
	var result int
	for _, b := range rules[node].requiredContents {
		result += b.count * (1+countBags(rules, b.bag))
	}
	return result
}