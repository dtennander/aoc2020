package main

import "math"

func sum(vals []int64) (sum int64) {
	for _, v := range vals {
		sum += v
	}
	return sum
}

func min(vals []int64) (r int64) {
	r = math.MaxInt64
	for _, v := range vals {
		if v < r {
			r = v
		}
	}
	return r
}

func max(vals []int64) (r int64) {
	for _, v := range vals {
		if v > r {
			r = v
		}
	}
	return r
}
