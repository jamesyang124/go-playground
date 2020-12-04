package threesum

import "sort"

// sort to reduce time, then comparison from both side
// threeSum ...
// only works when no duplicate elements
func threeSumNoDuplicate(nums []int) [][]int {
	target := 0
	// O(nlogn)
	sort.Ints(nums)
	output := [][]int{}

	// -2 since we don't have to iterate last 2 elements, it will check by inner loop
	// iterations are n-1 + n-2 + .... 2 => n*(n-1) / 2 => O(n^2)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			leftRightResult := nums[left] + nums[right]
			sum := nums[i] + leftRightResult

			if sum == target {
				output = append(output, []int{nums[i], nums[left], nums[right]})
				left++
				right--
			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return output
}
