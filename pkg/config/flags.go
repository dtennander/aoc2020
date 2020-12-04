package config

import (
	"aod2020/pkg/input"
	"flag"
	"io"
)

// ParseDefaultFlags parses the program arguments and returns exercise input and part number.
func ParseDefaultFlags() (part int, inputR io.Reader) {
	partP := flag.Int("part", 1, "The part of the assignment to run.")
	inputFile := flag.String("input", "stdin", "file path of the input file.")
	flag.Parse()
	reader, err := input.GetReader(*inputFile)
	if err != nil {
		panic(err.Error())
	}
	return *partP, reader
}
