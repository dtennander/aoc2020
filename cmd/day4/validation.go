package main

import (
	"regexp"
	"strconv"
)

func IsValidPassport(pairs map[string]string, validateValue bool) bool {
	for _, ef := range requiredFields {
		if v, ok := pairs[ef.key]; !ok || (validateValue && !ef.validation(v)) {
			return false
		}
		delete(pairs, ef.key)
	}
	delete(pairs, "cid")
	return len(pairs) == 0
}

type fieldRule struct {
	key string
	validation func(value string) bool
}

var requiredFields = []fieldRule{
	{"byr", func(v string) bool {
		year, err := strconv.ParseInt(v, 10, 64)
		return err == nil && 1920 <= year && year <= 2002
	}},
	{"iyr", func(v string) bool {
		year, err := strconv.ParseInt(v, 10, 64)
		return err == nil && 2010 <= year && year <= 2020
	}},
	{"eyr", func(v string) bool {
		year, err := strconv.ParseInt(v, 10, 64)
		return err == nil && 2020 <= year && year <= 2030
	}},
	{"hgt", func(v string) bool {
		val, err := strconv.ParseInt(v[:len(v)-2], 10, 64)
		unit := v[len(v)-2:]
		return err == nil && (
			unit == "cm" && 150 <= val && val <= 193 ||
			unit == "in" && 59 <= val && val <= 76)
	}},
	{"hcl", func(v string) bool {
		return validHexColorPattern.MatchString(v)
	}},
	{"ecl", func(v string) bool {
		return validEyeColors[v]
	}},
	{"pid", func(v string) bool {
		_, err := strconv.ParseInt(v, 10, 64)
		return err == nil && len(v) == 9
	}},
}

var validEyeColors = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

var validHexColorPattern = regexp.MustCompile("#[0-9a-f]")