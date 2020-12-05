package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestReadSeats(t *testing.T) {
	ids, lowest, highest := readSeatIds(bytes.NewBufferString("FFFFFFFRRR\nBBBBBBBRRL"))
	if lowest != 7 {
		t.Error("Didn't catch lowest id...")
	}
	if highest != 1022 {
		t.Error("Didn't catch highest id...")
	}
	for seat, status := range ids {
		shouldBeTaken := seat == 7 || seat == 1022
		if shouldBeTaken && status == Available {
			t.Error("seat 8 and 1022 are marked as available...")
		}
		if !shouldBeTaken && status == Taken {
			t.Error("seat is falsely marked as taken.")
		}
	}
}

func BenchmarkReadSeats(b *testing.B) {
	data, err := ioutil.ReadFile("../../input/day5.txt")
	if err != nil {
		b.Fail()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		readSeatIds(bytes.NewReader(data))
	}
}

func BenchmarkAvailableSeat(b *testing.B) {
	data, err := ioutil.ReadFile("../../input/day5.txt")
	if err != nil {
		b.Fail()
	}
	seats, lowestId, highestId := readSeatIds(bytes.NewReader(data))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		availableSeat(seats, lowestId, highestId)
	}
}
