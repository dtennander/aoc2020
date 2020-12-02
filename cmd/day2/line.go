package main

import (
	"regexp"
	"strconv"
	"strings"
)

type Line struct {
	letter   rune
	min      int64
	max      int64
	password string
}

var lineExpr = regexp.MustCompile("([0-9]+)-([0-9]+) (.): (.+)")

func NewLine(line string) (Line, error) {
	match := lineExpr.FindStringSubmatch(line)
	min, err := strconv.ParseInt(match[1], 10, 64)
	if err != nil {
		return Line{}, err
	}
	max, err := strconv.ParseInt(match[2], 10, 64)
	if err != nil {
		return Line{}, err
	}
	letter := []rune(match[3])[0]
	return Line{
		min:      min,
		max:      max,
		letter:   letter,
		password: match[4],
	}, nil
}

func (l Line) FollowsCountRule() bool {
	c := int64(strings.Count(l.password, string(l.letter)))
	return l.min <= c && c <= l.max
}

func (l Line) FollowsPositionRule() bool {
	arr := []rune(l.password)
	fst := arr[l.min-1]
	snd := arr[l.max-1]
	return (fst == l.letter) != (snd == l.letter)
}
