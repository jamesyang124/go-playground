package basic

// should always have respond type
func declaration(x, y int) int {
	return x + y
}

// ReturnMultiple (x int, y int)
func ReturnMultiple(x, y int) (int, int) {
	return x + y, x * y
}

// ReturnNamed (x int, y int)
// avoid use it for readability
func ReturnNamed(x, y int) (add, mul int) {
	add = x + y
	mul = x * y

	// If the result parameters are named, a return statement without arguments returns the current values of the results.
	return
}
