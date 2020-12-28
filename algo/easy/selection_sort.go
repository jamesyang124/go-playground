package easy

// SelectionSort ...
// find min put it to begin iteratively
// https://www.geeksforgeeks.org/selection-sort/
// T O(n^2)
// S O(1)
func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		s, tmp := i, array[i]
		for j := i + 1; j < len(array); j++ {
			if tmp > array[j] {
				s = j
				tmp = array[j]
			}
		}
		array[i], array[s] = array[s], array[i]
	}
	return array
}
