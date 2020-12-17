package powerset

// PowersetItr ...
// given array, populate its elements as power set
// ex: [1, 2, 3]
// as: [[], [1], [2], [1, 2], [3], [1, 3], [2, 3], [1, 2, 3]]
// T O(n*2^n)
// S O(n*2^n) in last run create array with n*2^n size though it used as output
func PowersetItr(array []int) [][]int {
	output := [][]int{{}}
	outputLen := len(output)

	for i := 0; i < len(array); i++ {
		// use count to avoid cost of re-calculate length of array
		j := 0
		for ; j < outputLen; j++ {
			// should deep copy instead of reuse same array reference
			elem := append([]int{}, output[j]...)
			elem = append(elem, array[i])
			output = append(output, elem)
		}
		outputLen += j
	}

	return output
}
