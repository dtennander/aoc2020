package main

import (
	"aod2020/pkg/config"
	"aod2020/pkg/input"
)

func main() {
	part, in := config.ParseDefaultFlags()
	numbers, err := input.ReadNumbers(in)
	if err != nil {
		println("Failed to parse input.")
		return
	}
	switch part {
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
