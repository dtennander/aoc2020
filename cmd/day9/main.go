package main

import (
	"aod2020/pkg/config"
	"aod2020/pkg/input"
)

func main() {
	part, in := config.ParseDefaultFlags()
	numbers, err := input.ReadNumbers(in)
	if err != nil {
		panic(err.Error())
	}
	invNumber := findInvalidNumber(numbers)
	println("Found number not matching:", invNumber)
	if part == 2 {
		summers := findContiguousSetThatSums(numbers, invNumber)
		println("Sum of min and max are:", min(summers) + max(summers))
	}
}

func findInvalidNumber(numbers []int64) int64 {
	for pointer := 25; pointer < len(numbers); pointer++ {
		if !sumExist(numbers[(pointer-25):pointer], numbers[pointer]) {
			return numbers[pointer]
		}
	}
	return 0
}

func sumExist(values []int64, sum int64) bool {
	for _, i := range values {
		for _, j := range values {
			if i+j == sum {
				return true
			}
		}
	}
	return false
}

func findContiguousSetThatSums(numbers []int64, target int64) []int64 {
	start := 0
	stop := 0
	for stop < len(numbers) {
		currentSum := sum(numbers[start:stop])
		if currentSum == target {
			return numbers[start:stop]
		}
		if currentSum < target {
			stop += 1
		} else if target < currentSum {
			start += 1
		}
	}
	return []int64{}
}