package main

import (
	"aod2020/pkg/config"
	"bufio"
	"io"
	"strconv"
)

func main() {
	part, in := config.ParseDefaultFlags()
	voltages := parseVoltages(in)
	if part == 1 {
		ones, twos, threes := findSpaces(voltages)
		println("ones:", ones, "twos", twos, "threes:", threes)
		println("ones times threes:", ones * threes)
	} else {
		interations := recursiveSearch(voltages, 0)
		println("Permutations:", interations)
	}
}

func parseVoltages(in io.Reader) []bool {
	s := bufio.NewScanner(in)
	var voltages [256]bool
	var maxVoltage uint64
	for s.Scan() {
		i, err := strconv.ParseUint(s.Text(), 10, 8)
		if err != nil {
			panic(err.Error())
		}
		if i > maxVoltage {
			maxVoltage = i
		}
		voltages[i] = true
	}
	voltages[maxVoltage + 3] = true
	return voltages[:(maxVoltage + 4)]
}

func findSpaces(array []bool) (ones int, twos int, threes int) {
	var lastTrue int
	for i, exists := range array {
		if exists {
			switch i - lastTrue {
			case 1:
				ones += 1
			case 2:
				twos += 1
			case 3:
				threes += 1
			default:
				panic("FAILED!")
			}
			lastTrue = i
		}
	}
	return ones, twos, threes
}

var savedResults = make(map[int]int)
func recursiveSearch(arr []bool, i int) int {
	result, ok := savedResults[i]
	if ok {
		return result
	}
	if len(arr) <= 4 {
		return 1
	}
	if len(arr) > 2 && arr[1] {
		result += recursiveSearch(arr[1:], i + 1)
	}
	if len(arr) > 3 && arr[2] {
		result += recursiveSearch(arr[2:], i + 2)
	}
	if len(arr) > 4 && arr[3] {
		result += recursiveSearch(arr[3:], i + 3)
	}
	savedResults[i] = result
	return result
}
