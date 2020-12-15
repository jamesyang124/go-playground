package medium

// MoveElementToEnd ...
// https://www.algoexpert.io/questions/Move%20Element%20To%20End
// moves all instances of that integer in the array to the end of the array.
// [5, 1, 2, 5, 5, 3, 4, 6, 7, 5, 8, 9, 10, 11, 5, 5, 12], 5
// [12, 1, 2, 11, 10, 3, 4, 6, 7, 9, 8, 5, 5, 5, 5, 5, 5]
// T O(n)
// S O(1)
func MoveElementToEnd(ary []int, toMove int) []int {
	i, j := 0, len(ary)-1
	for i < j {
		for i < j && ary[j] == toMove {
			j--
		}

		if ary[i] == toMove {
			ary[i], ary[j] = ary[j], ary[i]
		}
		i++
	}

	return ary
}
