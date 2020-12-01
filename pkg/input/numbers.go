package input

import (
	"bufio"
	"io"
	"strconv"
)

// ReadNumberInput takes a reader, usually stdIn, and returns a list of all line separated numbers.
func ReadNumberInput(reader io.Reader) (result []int64, err error) {
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
