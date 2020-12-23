package easy

import "math"

// FindThreeLargest ...
// array at least have size 3, find 3 largest number in it
// sort if find nth largest, n*logn
//
// T O(n)
// S O(1)
func FindThreeLargest(array []int) []int {
	res := []int{math.MinInt64, math.MinInt64, math.MinInt64}

	for _, e := range array {
		if res[2] < e {
			res[0] = res[1]
			res[1] = res[2]
			res[2] = e
			continue
		}

		if res[1] < e {
			res[0] = res[1]
			res[1] = e
			continue
		}

		if res[0] < e {
			res[0] = e
		}
	}

	return res
}
