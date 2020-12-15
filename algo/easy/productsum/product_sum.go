package productsum

// https://www.algoexpert.io/questions/Product%20Sum

// A "special" array is a non-empty array that contains either integers or other
// "special" arrays. The product sum of a "special" array is the sum of its
// elements, where "special" arrays inside it are summed themselves and then
// multiplied by their level of depth.
// The depth of a "special" array is how far nested it is.
// For instance, the depth of [[]] is 2, [[[]]] is 3
// [x, [y, [z]]] is x + 2 * ( y + 3 * z)
//
// [5, 2, [7, -1], 3, [6, [-13, 8], 4]]
// 12

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
// T O(n)
// S O(d), d < n
func ProductSum(array []interface{}) int {
	return recursion(array, 1)
}
