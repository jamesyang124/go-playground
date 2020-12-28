package easy

// InsertionSort ...
// T O()
// S O(1)
func InsertionSort(array []int) []int {

	// from begin
	/*
		for i := 1; i < len(array); i++ {
			tmp := array[i]
			for j := 0; j < i; j++ {
				if array[j] >= tmp {
					array[j], tmp = tmp, array[j]
				}
			}
			array[i] = tmp
		}
	*/

	// can short circuit if before < current
	for i := 1; i < len(array); i++ {
		for j := i - 1; j >= 0; j-- {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			} else {
				break
			}
		}
	}

	return array
}
