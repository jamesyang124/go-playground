package threesum

import "sort"

// leetcode 15
// https://leetcode.com/problems/3sum/
// given array may have duplicate elements, output unique triplets
// [-2,0,1,1,2]
// => [[-2,0,2],[-2,1,1]]
//
// [-4,-2,-2,-2,0,1,2,2,2,3,3,4,4,6,6]
// => [[-4,-2,6],[-4,0,4],[-4,1,3],[-4,2,2],[-2,-2,4],[-2,0,2]]
//
// [-2,0,0,2,2]
// => [[-2,0,2]]

// ThreeSum ...
// sort to reduce time, then comparison from both side
// works when duplicate elements
// faster than other 95%
// 28 ms
// usually O(n^(k-1))
// T O(n^2)
// S O(1)
func ThreeSum(nums []int) [][]int {
	target := 0
	// O(nlogn)
	sort.Ints(nums)
	output := [][]int{}

	// -2 since we don't have to iterate last 2 elements, it will check by inner loop
	// iterations are n-1 + n-2 + .... 2 => n*(n-1) / 2 => O(n^2)
	for i := 0; i < len(nums)-2; i++ {
		left, right := i+1, len(nums)-1

		// handle duplicate elements
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for left < right {
			leftRightResult := nums[left] + nums[right]
			sum := nums[i] + leftRightResult

			if sum == target {
				output = append(output, []int{nums[i], nums[left], nums[right]})
				left++
				right--

				// handle duplicate elements,
				// same combination is handled in this iteration
				for left < right && nums[left] == nums[left-1] {
					left++
				}

				// handle duplicate elements,
				// same combination is handled in this iteration
				for left < right && nums[right] == nums[right+1] {
					right--
				}

			} else if sum > target {
				right--
			} else {
				left++
			}
		}
	}

	return output
}
