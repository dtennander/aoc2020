package main

import (
	"aod2020/pkg/config"
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type PlaneSeats [1024]SeatStatus
type SeatId uint16

type SeatStatus bool
const (
	Taken SeatStatus = true
	Available SeatStatus = false
)


func main() {
	_, r := config.ParseDefaultFlags()
	seats, lowestId, highestId := readSeatIds(r)
	fmt.Printf("Highest Id taken is %v\n", highestId)
	fmt.Printf("Lowest Id taken is %v\n", lowestId)
	ourSeat := availableSeat(seats, lowestId, highestId)
	fmt.Printf("Your seat is %v\n", ourSeat)
}

func readSeatIds(r io.Reader) (seats PlaneSeats, lowest SeatId, highest SeatId) {
	s := bufio.NewScanner(r)
	lowest = 1024
	for s.Scan() {
		id := parseSeatId(s.Text())
		seats[id] = Taken
		if id > highest {
			highest = id
		}
		if id < lowest {
			lowest = id
		}
	}
	return seats, lowest, highest
}

var repl = strings.NewReplacer("B", "1", "R", "1", "F", "0", "L", "0")
func parseSeatId(line string) SeatId {
	id, err := strconv.ParseUint(repl.Replace(line), 2, 16)
	if err != nil {
		panic("Could not parse line: \"" + line + "\"")
	}
	return SeatId(id)
}

func availableSeat(seats PlaneSeats, start SeatId, stop SeatId) SeatId {
	for i, seat := range seats[start:stop] {
		if seat == Available {
			return start +  SeatId(i)
		}
	}
	return 0
}
