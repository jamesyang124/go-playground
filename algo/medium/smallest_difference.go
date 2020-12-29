package medium

import (
	"math"
	"sort"
)

// SmallestDifference ...
// https://www.algoexpert.io/questions/Smallest%20Difference
//
// sort first, then compare either cases:
//
// a0 a1 b0 b1
// b0 b1 a0 a1
// a0 b0 a1 b1
// a0 b0 b1 a1
// b0 a0 a1 b1
// b0 a0 b1 a1
//
// if ai > bj, j + 1
// if bj > ai, i + 1
// then compare abs diff again
//
// T O (nlogn + mlogm)
// S O (1)
func SmallestDifference(array1, array2 []int) []int {
	var pair []int

	sort.Ints(array1)
	sort.Ints(array2)

	current, smallest := math.MaxInt32, math.MaxInt32
	aIdx, bIdx := 0, 0

	for aIdx < len(array1) && bIdx < len(array2) {
		a, b := array1[aIdx], array2[bIdx]
		if a > b {
			current = a - b
			bIdx++
		} else if b > a {
			current = b - a
			aIdx++
		} else {
			return []int{a, b}
		}

		if smallest > current {
			smallest = current
			pair = []int{a, b}
		}
	}

	return pair
}
