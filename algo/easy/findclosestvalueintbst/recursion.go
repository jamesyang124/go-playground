package findclosestvalueinbst

func absdiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func recursion(tree *BST, target, closest int) int {
	if absdiff(closest, target) > absdiff(tree.Value, target) {
		closest = tree.Value
	}

	if tree.Right != nil && tree.Value < target {
		return recursion(tree.Right, target, closest)
	} else if tree.Left != nil && tree.Value > target {
		return recursion(tree.Left, target, closest)
	}

	return closest
}

// Recursion ...
// https://www.algoexpert.io/questions/Find%20Closest%20Value%20In%20BST
// assume tree at least have 1 node
//     10
//   5    15
//      12
// target 13
// otuput 12
//
// T O(logn)
// S O(logn)
func (tree *BST) Recursion(target int) int {
	return recursion(tree, target, tree.Value)
}
