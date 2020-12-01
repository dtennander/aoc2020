package input

import (
	"io"
	"os"
)

// GetReader returns the reader to use for the input.
// If the input file is `stdin` os.StdIn will be used as the source.
func GetReader(inputFile string) (io.Reader, error) {
	if inputFile == "stdin" {
		return os.Stdin, nil
	}
	return os.Open(inputFile)
}
