package binarysearch

// Iteration ...
// T O(logn)
// S O(1)
func Iteration(array []int, target int) int {
	res := -1
	l, r := 0, len(array)-1

	for l <= r && l >= 0 && r < len(array) {
		mid := (l + r) / 2
		if array[mid] == target {
			res = mid
			break
		} else if array[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return res
}
