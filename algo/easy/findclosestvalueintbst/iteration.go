package findclosestvalueinbst

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

// Iteration ...
// https://www.algoexpert.io/questions/Find%20Closest%20Value%20In%20BST
// T O(n), avg O(logn)
// S O(1)
func (tree *BST) Iteration(target int) int {
	var output, diff int

	if tree != nil {
		output = tree.Value
		diff = absdiff(tree.Value, target)
	}

	for tree != nil {
		if tree.Left != nil && diff > absdiff(tree.Left.Value, target) {
			diff = absdiff(tree.Left.Value, target)
			output = tree.Left.Value
		}

		if tree.Right != nil && diff > absdiff(tree.Right.Value, target) {
			diff = absdiff(tree.Right.Value, target)
			output = tree.Right.Value
		}

		if tree.Value < target {
			tree = tree.Right
		} else {
			tree = tree.Left
		}
	}

	return output
}
