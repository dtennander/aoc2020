package main

import (
	"aod2020/pkg/config"
	"bufio"
	"io"
	"math"
)

func main() {
	p, in := config.ParseDefaultFlags()
	var (
		sightDistance int
		tolerance     int
	)
	if p == 1 {
		sightDistance, tolerance = 1, 4
	} else {
		sightDistance, tolerance = math.MaxInt64, 5
	}
	a := parseArea(in)
	changed := true
	for changed {
		a, changed = a.step(tolerance, sightDistance)
	}
	println()
	println(a.String())
	println("Number of occupied seats:", a.countOccupied())
}


type Area [][]Tile
type Tile rune
const (
	Empty Tile = 'L'
	Floor Tile = '.'
	Occupied Tile = '#'
)

func parseArea(in io.Reader) (area Area) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		area = append(area, []Tile(s.Text()))
	}
	return area
}

func (a Area) DeepCopy() Area {
	nextArea := make(Area, len(a))
	for i := range a {
		nextArea[i] = make([]Tile, len(a[i]))
		copy(nextArea[i], a[i])
	}
	return nextArea
}

func (a Area) step(tolerance, sightDistance int) (Area, bool) {
	changed := false
	nextArea := a.DeepCopy()
	for i, l := range a {
		for j, t := range l {
			if t == Floor {
				continue
			}
			occupied := a.countSeen(i, j, sightDistance)
			switch {
			case t == Empty && occupied == 0:
				nextArea[i][j] = Occupied
				changed = true
			case t == Occupied && occupied >= tolerance:
				nextArea[i][j] = Empty
				changed = true
			}
		}
	}
	return nextArea, changed
}

var directions = [8]Vector{
	{-1, 1},{0, 1},{1, 1},
	{-1, 0},       {1, 0},
	{-1,-1},{0,-1},{1,-1},
}
func (a Area) countSeen(i, j, maxSteps int) (count int) {
	for _, d := range directions {
		steps:
		for s := 1; s <= maxSteps; s++ {
			k,l := d.Multi(s).Plus(i,j)
			switch  {
			case k < 0, l < 0, len(a) <= k, len(a[i]) <= l:
				break steps
			case a[k][l] == Floor:
				continue steps
			case a[k][l] == Empty:
				break steps
			case a[k][l] == Occupied:
				count++
				break steps
			}
		}
	}
	return count
}

func (a Area) countOccupied() (occupied int) {
	for i  := range a {
		for j := range a[i] {
			if a[i][j] == Occupied {
				occupied++
			}
		}
	}
	return occupied
}

func (a Area) String() (str string) {
	for _, l := range a {
		for _, t := range l {
			str += string(t)
		}
		str += "\n"
	}
	return str
}
