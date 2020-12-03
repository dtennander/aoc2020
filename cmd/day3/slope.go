package main

import (
	"bufio"
	"io"
)

type Patch int
const (
	Open Patch = iota
	Tree
)

type Slope [][]Patch

func (s Slope) Get(i, j int) Patch {
	n := len(s)
	m := len(s[0])
	return s[i % n][j % m]
}

func (s Slope) Length() int {
	return len(s)
}

func ReadSlope(r io.Reader) Slope {
	scanner := bufio.NewScanner(r)
	var patch Slope
	for scanner.Scan() {
		textLine := scanner.Text()
		var line []Patch
		for _, c := range textLine {
			switch c {
			case '.':
				line = append(line, Open)
			case '#':
				line = append(line, Tree)
			}
		}
		patch = append(patch, line)
	}
	return patch
}

