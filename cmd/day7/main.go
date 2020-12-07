package main

import (
	"aod2020/pkg/config"
	"bufio"
	"io"
	"strconv"
	"strings"
)

func main() {
	part, in := config.ParseDefaultFlags()
	rules := readRules(in)
	if part == 1 {
		println("The number of allowed bags are:", walkConnectedRules(rules, "shiny gold", make(map[string]bool)))
	}
	if part == 2 {
		println("The golden bag contains:", countBags(rules, "shiny gold"))
	}
}

type Rule struct {
	bag     string
	allowed []BagCount
}

type BagCount struct {
	count int
	bag string
}

func readRules(in io.Reader) (result map[string]Rule) {
	s := bufio.NewScanner(in)
	result = make(map[string]Rule)
	for s.Scan() {
		words := strings.Fields(s.Text())
		name := words[0] + " " + words[1]
		result[name] = Rule{
			bag:     name,
			allowed: parseBagCounts(words[4:]),
		}
	}
	return result
}

func parseBagCounts(rules []string) (containedBags []BagCount) {
	if rules[0] == "no" { // "No bag count" - Leaf.
		return
	}
	for len(rules) != 0 {
		count, err := strconv.ParseInt(rules[0], 10, 64)
		if err != nil {
			panic(err.Error())
		}
		containedBags = append(containedBags, BagCount{
			count: int(count),
			bag: rules[1] + " " + rules[2],
		})
		rules = rules[4:]
	}
	return containedBags
}

func walkConnectedRules(rules map[string]Rule, node string, visited map[string]bool) int {
	for _, r := range rules {
		for _, a := range r.allowed {
			if a.bag == node && !visited[r.bag]{
				visited[r.bag] = true
				walkConnectedRules(rules, r.bag, visited)
			}
		}
	}
	return len(visited)
}

func countBags(rules map[string]Rule, node string) int {
	var result int
	for _, b := range rules[node].allowed {
		result += b.count * (1+countBags(rules, b.bag))
	}
	return result
}