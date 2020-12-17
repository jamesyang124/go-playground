package powerset

func helper(array []int, index int) [][]int {
	var res [][]int
	if index < 0 {
		res = [][]int{{}}
	}

	// ....

	return res
}

// Recursion ...
func Recursion(array []int) [][]int {
	helper(array, len(array)-1)
}
