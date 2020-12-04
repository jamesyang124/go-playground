package threesum

import (
	"sort"
)

// faster than other 38%
// waste time to create usedMap map and its look-up time
// 292 ms
func threeSumSlow(nums []int) [][]int {
	target := 0
	// O(nlogn)
	sort.Ints(nums)
	output := [][]int{}

	usedMap := make(map[int]map[int]int)

	// -2 since we don't have to iterate last 2 elements, it will check by inner loop
	// iterations are n-1 + n-2 + .... 2 => n*(n-1) / 2 => O(n^2)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		for left < right {

			leftRightResult := nums[left] + nums[right]
			sum := nums[i] + leftRightResult

			_, leftUsed := usedMap[nums[i]][nums[left]]
			_, rightUsed := usedMap[nums[i]][nums[right]]

			if leftUsed == true && rightUsed == true {
				left++
				right--
				continue
			}

			if sum == target {
				output = append(output, []int{nums[i], nums[left], nums[right]})

				_, ok := usedMap[nums[i]]

				if ok == false {
					usedMap[nums[i]] = make(map[int]int)
				}

				usedMap[nums[i]][nums[left]] = nums[left]
				usedMap[nums[i]][nums[right]] = nums[right]

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
