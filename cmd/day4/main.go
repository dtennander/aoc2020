package main

import (
	"aod2020/pkg/config"
	"bufio"
	"bytes"
	"fmt"
	"strings"
)

func main() {
	part, in := config.ParseDefaultFlags()
	answer := 0
	s := bufio.NewScanner(in)
	s.Split(scanEntries)
	for s.Scan() {
		t := s.Text()
		pairs := make(map[string]string)
		for _, f := range strings.Fields(t) {
			pair := strings.Split(f, ":")
			pairs[pair[0]] = pair[1]
		}
		if IsValidPassport(pairs, part == 2) {
			answer += 1
		}
	}
	fmt.Printf("Number of Valid passports: %v", answer)
}

// scanEntries Returns bytes for each Entry in the batch file.
func scanEntries(data []byte, atEOF bool) (advance int, token []byte, err error) {
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

