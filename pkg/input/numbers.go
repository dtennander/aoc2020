package input

import (
	"bufio"
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
