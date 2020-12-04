package main

import (
	"aod2020/pkg/config"
	"aod2020/pkg/input"
	"bufio"
	"fmt"
	"strings"
)

func main() {
	part, in := config.ParseDefaultFlags()
	answer := 0
	s := bufio.NewScanner(in)
	s.Split(input.SplitEmptyLine)
	for s.Scan() {
		t := s.Text()
		pairs := make(map[string]string)
		for _, f := range strings.Fields(t) {
			pair := strings.Split(f, ":")
			pairs[pair[0]] = pair[1]
		}
		if IsValidPassport(pairs, part == 2) {
			answer += 1
		}
	}
	fmt.Printf("Number of Valid passports: %v", answer)
}

