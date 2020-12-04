package productsum

// https://www.algoexpert.io/questions/Product%20Sum

// SpecialArray ...
type SpecialArray []interface{}

// O(n) time, n is the total numbers (including sub-elements) in the array
// O(d) space, d is the greatest depth of special arrays in the array.

// https://stackoverflow.com/questions/43298938/space-complexity-of-recursive-function
// Our memory complexity is determined by the number of return statements
// because each function call will be stored on the program stack.
// To generalize, a recursive function's memory complexity is O(recursion depth).
func recursion(array SpecialArray, mult int) int {
	sum := 0
	for _, e := range array {
		if cast, ok := e.(SpecialArray); ok {
			sum += recursion(cast, mult+1)
		} else if cast, ok := e.(int); ok {
			sum += cast
		}
	}
	return sum * mult
}

// ProductSum ...
// Tip: Each element of `array` can either be cast
// to `SpecialArray` or `int`.
func ProductSum(array []interface{}) int {
	return recursion(array, 1)
}
