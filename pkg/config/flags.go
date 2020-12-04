package config

import (
	"flag"
	"io"
	"os"
)

// ParseDefaultFlags parses the program arguments and returns exercise input and part number.
func ParseDefaultFlags() (part int, inputR io.Reader) {
	partP := flag.Int("part", 1, "The part of the assignment to run.")
	inputFile := flag.String("input", "stdin", "file path of the input file.")
	flag.Parse()
	reader, err := getReader(*inputFile)
	if err != nil {
		panic(err.Error())
	}
	return *partP, reader
}

// getReader returns the reader to use for the input.
// If the input file is `stdin` os.StdIn will be used as the source.
func getReader(inputFile string) (io.Reader, error) {
	if inputFile == "stdin" {
		return os.Stdin, nil
	}
	return os.Open(inputFile)
}
