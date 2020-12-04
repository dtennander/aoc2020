package main

import (
	"aod2020/pkg/config"
	"bufio"
	"fmt"
	"io"
)

func main() {
	part, in := config.ParseDefaultFlags()
	switch part {
	case 1:
		scanAndCountValidLines(in, Line.FollowsCountRule)
	case 2:
		scanAndCountValidLines(in, Line.FollowsPositionRule)
	default:
		println("Only parts 1 or 2 are available.")
	}

}

func scanAndCountValidLines(in io.Reader, lineValidator func(Line) bool) {
	scanner := bufio.NewScanner(in)
	validLines := 0
	for scanner.Scan() {
		line, err := NewLine(scanner.Text())
		if err != nil {
			panic(err.Error())
		}
		if lineValidator(line) {
			validLines += 1
		}
	}
	fmt.Printf("Valid entries: %v", validLines)
}
