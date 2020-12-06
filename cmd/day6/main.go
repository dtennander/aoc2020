package main

import (
	"aod2020/pkg/config"
	"aod2020/pkg/input"
	"bufio"
	"fmt"
	"math/bits"
	"strings"
)

func main() {
	part, r := config.ParseDefaultFlags()
	s := bufio.NewScanner(r)
	s.Split(input.SplitEmptyLine)
	sum := 0
	for s.Scan() {
		group := s.Text()
		if part == 1 {
			sum += getUnionSize(group)
		} else {
			sum += getIntersectionSize(group)
		}
	}
	if part == 1 {
		fmt.Println("The union of the yes answers:", sum)
	} else {
		fmt.Println("The intersection of the yes answers:", sum)
	}
}

// AllYes is the answer yes on all 26 questions.
const AllYes uint32 = (1 << 26) - 1
// AllNo is the answer no on all 26 questions.
const AllNo uint32 = 0

func getUnionSize(group string) int {
	result := AllNo
	for _, f := range strings.Fields(group) {
		result |= parseAsBits(f)
	}
	return bits.OnesCount32(result)
}

func getIntersectionSize(group string) int {
	result := AllYes
	for _, f := range strings.Fields(group) {
		result &= parseAsBits(f)
	}
	return bits.OnesCount32(result)
}

// parseAsBits takes a line of answers and returns a 26Bit representation of those answers.
func parseAsBits(line string) (answer uint32) {
	for _, r := range line {
		answer |= 1 << (r - 'a')
	}
	return answer
}
