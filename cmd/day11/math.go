package main

type Vector [2]int

func (v Vector) Multi(n int) Vector {
	return Vector{n*v[0], n*v[1]}
}

func (v Vector) Plus(i, j int) (int, int) {
	return i + v[0], j + v[1]
}
