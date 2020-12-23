package easy

// BubbleSort ...
// could naively iterate n^2 times,
// keep swapping comparison with next element, iterate n times for whole array scan
// but we could shortcut when no longer have to swapped
//
// T O(n^2)
// S O(n)
func BubbleSort(array []int) []int {
	swapped := true

	for swapped == true {
		swapped = false
		for j := 0; j < len(array)-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				swapped = true
			}
		}
	}

	return array
}
