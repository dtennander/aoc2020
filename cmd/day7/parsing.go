package main

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

func readRules(in io.Reader) (result map[string]Rule) {
	s := bufio.NewScanner(in)
	result = make(map[string]Rule)
	for s.Scan() {
		words := strings.Fields(s.Text())
		name := words[0] + " " + words[1]
		result[name] = Rule{
			bag:              name,
			requiredContents: parseBagCounts(words[4:]),
		}
	}
	return result
}

// Type alias to reduce boiler plate
type bagCount = struct {
	bag string
	count int
}

func parseBagCounts(rules []string) (containedBags []bagCount) {
	if rules[0] == "no" { // "No bag count" - Leaf.
		return
	}
	for len(rules) != 0 {
		count, err := strconv.ParseInt(rules[0], 10, 64)
		if err != nil {
			panic(err.Error())
		}
		containedBags = append(containedBags, bagCount{
			bag: rules[1] + " " + rules[2],
			count: int(count),
		})
		rules = rules[4:]
	}
	return containedBags
}
