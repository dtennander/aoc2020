package main

import (
	"aod2020/pkg/input"
	"bufio"
	"flag"
	"fmt"
	"io"
)

var part = flag.Int("part",1, "The part of the assignment to run.")
var inputFile = flag.String("input", "stdin", "file path of the input file.")

func main() {
	flag.Parse()
	in, err := input.GetReader(*inputFile)
	if err != nil { println("Failed to open input."); return }
	switch *part {
	case 1:
		scanAndCountValidLines(in, Line.CountRuleValid)
	case 2:
		scanAndCountValidLines(in, Line.PositionRuleValid)
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