package medium

// MoveElementToFront ...
// [5, 1, 2, 5, 5, 3, 4, 6, 7, 5, 8, 9, 10, 11, 5, 5, 12], 5
//
// T O(n)
// S O(1)
func MoveElementToFront(ary []int, toMove int) []int {
	i, j := 0, len(ary)-1

	for i < j {
		for i < j && ary[i] == toMove {
			i++
		}

		if ary[j] == toMove {
			ary[i], ary[j] = ary[j], ary[i]
		}
		j--
	}
	return ary
}
