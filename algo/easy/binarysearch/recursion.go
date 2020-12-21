package binarysearch

func helper(sorted []int, i, j, target int) int {
	if i > j {
		return -1
	}

	mid := (i + j) / 2

	if sorted[mid] == target {
		return mid
	}

	if sorted[mid] > target {
		return helper(sorted, i, mid-1, target)
	}

	return helper(sorted, mid+1, j, target)
}

// Recursion ...
func Recursion(sorted []int, target int) int {
	return helper(sorted, 0, len(sorted)-1, target)
}
