package findclosestvalueinbst

// Iteration ...
// https://www.algoexpert.io/questions/Find%20Closest%20Value%20In%20BST
//
// check absolute diff with previous closest and current node
// keep iterate throguh as binary searching to ensure not missing
// indirect nodes
//
// T O(n), avg O(logn)
// S O(1)
func (tree *BST) Iteration(target int) int {
	var closest int

	if tree != nil {
		closest = tree.Value
	}

	for tree != nil {
		if absdiff(closest, target) > absdiff(tree.Value, target) {
			closest = tree.Value
		}

		if tree.Value < target {
			tree = tree.Right
		} else {
			tree = tree.Left
		}
	}

	return closest
}
