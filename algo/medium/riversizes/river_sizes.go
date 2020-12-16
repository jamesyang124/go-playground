package riversizes

func recursion(matrix [][]int, i, j int) int {
	if matrix[i][j] == 0 {
		return 0
	}

	// in-place trick, so we don't need extra visisted matrix to record visited node
	matrix[i][j] = 0
	sum := 1

	if i > 0 {
		sum += recursion(matrix, i-1, j)
	}
	if i < len(matrix)-1 {
		sum += recursion(matrix, i+1, j)
	}

	if j > 0 {
		sum += recursion(matrix, i, j-1)
	}

	if j < len(matrix[i])-1 {
		sum += recursion(matrix, i, j+1)
	}

	return sum
}

// RiverSizes ...
// https://www.algoexpert.io/questions/River%20Sizes
// 2D array, 0 => land, 1 => part of river
// output each possible river length, in arbitrary order
//
// [
//   [1, 0, 0, 1, 0],
//   [1, 0, 1, 0, 0],
//   [0, 0, 1, 0, 1],
//   [1, 0, 1, 0, 1],
//   [1, 0, 1, 1, 0]
// ]
// [2, 1, 5, 2, 2]
//
// T O(n), n is element size w * h (width * height)
// S O(n), worst case is recursively circular river, which is ~= n/2 size ex:
//         [
//           [1, 1, 1, 1, 1, 1],
//           [0, 0, 0, 0, 0, 1],
//           [1, 1, 1, 1, 0, 1],
//           [1, 0, 0, 1, 0, 1],
//           [1, 0, 0, 0, 0, 1],
//           [1, 1, 1, 1, 0, 1]
//         ]
//         if inplace-replacement is not allowed, S would be O(n) as visited matrix w * h
func RiverSizes(matrix [][]int) []int {
	res := []int{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			riverLen := recursion(matrix, i, j)

			if riverLen > 0 {
				res = append(res, riverLen)
			}
		}
	}

	return res
}
