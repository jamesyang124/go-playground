package medium

// ArrayOfProducts ...
// https://www.algoexpert.io/questions/Array%20Of%20Products
// [5, 1, 4, 2]
// [8, 40, 10, 20]
// 8 = 1 * 4 * 2
//
// T O(n^2) => n - 1 + n -2 + .... 1
// S O(1) => 1
func ArrayOfProducts(array []int) []int {
	output := []int{}
	productBase := 1

	for i := 0; i < len(array); i++ {
		products := 1
		for j := i + 1; j < len(array); j++ {
			products = products * array[j]
		}
		output = append(output, productBase*products)

		productBase = productBase * array[i]
	}

	return output
}
