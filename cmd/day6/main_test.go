package main

import (
	"strings"
	"testing"
)

const data = "abc\nabc\nabx"

func BenchmarkGetIntersectionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getIntersectionSize(data)
	}
}

func BenchmarkGetUnionSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getUnionSize(data)
	}
}

// getUnionSizeMap Is the naive implementation of getUnionSize.
// It is here so we can compare the benchmarks.
func getUnionSizeMap(group string) int {
	hash := make(map[rune]bool)
	for _, f := range strings.Fields(group) {
		for _, r := range f {
			hash[r] = true
		}
	}
	return len(hash)
}

func BenchmarkGetUnionSizeMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		getUnionSizeMap(data)
	}
}