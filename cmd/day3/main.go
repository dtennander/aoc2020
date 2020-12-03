package main

import (
	"aod2020/pkg/input"
	"flag"
	"fmt"
)

var alternatives = [][][]int{
	{
		// Part 1
		{1, 3},
	},{
		// Part 2
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	},
}

var part = flag.Int("part", 1, "The part of the assignment to run.")
var inputFile = flag.String("input", "stdin", "file path of the input file.")

func main() {
	flag.Parse()
	r, err := input.GetReader(*inputFile)
	if err != nil {
		panic(err.Error())
	}
	slope := ReadSlope(r)
	trees := 1
	for _, alt := range alternatives[*part - 1] {
		trees *= countTrees(slope, alt[0], alt[1])
	}
	fmt.Printf("Hit trees: %v", trees)
}

func countTrees(slope Slope, downStep, rightStep int) int {
	trees := 0
	for i, j := 0, 0; i < slope.Length(); i, j = i+downStep, j+rightStep {
		if slope.Get(i, j) == Tree {
			trees += 1
		}
	}
	return trees
}