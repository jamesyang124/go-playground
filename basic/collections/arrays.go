package collections

// Go only allows numbers & strings to be const'ed. <- ?
// const (
// 	IntArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	// [...] as implicit length
// 	ImplicitLengthArray = [...]string{"hello", "it's", "me"}
//
// 	// const must init with {}
// 	multiDimensionArray = [2][1]int{}
// )

var (
	multiDimensionArray = [2][1]int{}
	// ImplicitLengthArray ...
	ImplicitLengthArray = [...]string{"hello", "it's", "me"}
	// IntArray ...
	IntArray = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// MultiDimensionArray ...
	MultiDimensionArray [2][1]string
)
