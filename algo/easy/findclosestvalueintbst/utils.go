package findclosestvalueinbst

// BST binary search tree
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
