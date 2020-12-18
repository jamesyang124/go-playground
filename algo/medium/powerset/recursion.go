package powerset

func helper(array []int, index int) [][]int {
	if index < 0 {
		return [][]int{{}}
	}

	subsets := helper(array, index-1)
	res := [][]int{}
	// deep copy
	for j := range subsets {
		res = append(res, subsets[j])
	}

	for i := 0; i < len(subsets); i++ {
		set := append([]int{}, subsets[i]...)
		newsets := append(set, array[index])
		res = append(res, newsets)
	}
	return res
}

// Recursion ...
// T O(n*2^n) 1, 2, 4, 8, 16...
// S O(n*2^2) stack space use as n, but tmp array grow as large as n*2^n
func Recursion(array []int) [][]int {
	return helper(array, len(array)-1)
}
