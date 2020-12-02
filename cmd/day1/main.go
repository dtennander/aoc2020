package main

import (
	"aod2020/pkg/input"
	"flag"
)

var part = flag.Int("part", 1, "The part of the assignment to run.")
var inputFile = flag.String("input", "stdin", "file path of the input file.")

func main() {
	flag.Parse()
	in, err := input.GetReader(*inputFile)
	if err != nil {
		println("Failed to open input.")
		return
	}
	numbers, err := input.ReadNumbers(in)
	if err != nil {
		println("Failed to parse input.")
		return
	}
	switch *part {
	case 1:
		part1(numbers)
	case 2:
		part2(numbers)
	default:
		println("Only parts 1 or 2 are available.")
	}
}

func part1(input []int64) {
	for _, i := range input {
		for _, j := range input {
			if i+j == 2020 {
				println(i * j)
				return
			}
		}
	}
}

func part2(input []int64) {
	for _, i := range input {
		for _, j := range input {
			for _, k := range input {
				if i+j+k == 2020 {
					println(i * j * k)
					return
				}
			}
		}
	}
}
