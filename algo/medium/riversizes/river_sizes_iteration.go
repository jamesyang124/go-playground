package riversizes

func otherNextNodes(matrix [][]int, visited [][]bool, i, j int) [][]int {
	nextNodes := [][]int{}

	if i > 0 && !visited[i-1][j] {
		nextNodes = append(nextNodes, []int{i - 1, j})
	}

	if j > 0 && !visited[i][j-1] {
		nextNodes = append(nextNodes, []int{i, j - 1})
	}

	if i < len(matrix)-1 && !visited[i+1][j] {
		nextNodes = append(nextNodes, []int{i + 1, j})
	}

	if j < len(matrix[i])-1 && !visited[i][j+1] {
		nextNodes = append(nextNodes, []int{i, j + 1})
	}

	return nextNodes
}

func traverseNodes(matrix [][]int, visited [][]bool, i, j int, res []int) []int {
	riverSizes := 0
	nextNodes := [][]int{{i, j}}

	for len(nextNodes) > 0 {
		loc := nextNodes[0]
		nextNodes = nextNodes[1:]
		i, j := loc[0], loc[1]
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		if matrix[i][j] == 0 {
			continue
		}

		riverSizes++
		neighbors := otherNextNodes(matrix, visited, i, j)
		for _, n := range neighbors {
			nextNodes = append(nextNodes, n)
		}
	}

	if riverSizes > 0 {
		res = append(res, riverSizes)
	}
	return res
}

// https://www.algoexpert.io/questions/River%20Sizes
// T O(n), n is element size w * h (width * height)
// S O(n)
func riverSizesIteration(matrix [][]int) []int {
	res := []int{}
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if visited[i][j] {
				continue
			}
			res = traverseNodes(matrix, visited, i, j, res)
		}
	}

	return res
}
