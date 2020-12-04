package input

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
)

// ReadNumbers takes a reader, usually os.StdIn, and returns a list of all line-separated numbers.
func ReadNumbers(reader io.Reader) (result []int64, err error) {
	stdin := bufio.NewScanner(reader)
	for stdin.Scan() {
		i, err := strconv.ParseInt(stdin.Text(), 10, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

var _ bufio.SplitFunc = SplitEmptyLine

// SplitEmptyLine Returns bytes for each Entry in the batch file.
func SplitEmptyLine(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.Index(data, []byte("\n\n")); i >= 0 {
		// We have a full newline-terminated line.
		return i + 1, data[0:i], nil
	}
	// If we're at EOF, we have a final, non-terminated line. Return it.
	if atEOF {
		return len(data), data, nil
	}
	// Request more data.
	return 0, nil, nil
}
