package collections

// const aMap = map[int]string{1: "1", 2: "2"}
//
// go does not have const map
// https://groups.google.com/forum/#!topic/golang-nuts/xUaDMciyB5Q
// Go only allows numbers & strings to be const'ed. <- ?

var (
	// IntStringMap ...
	IntStringMap = map[int]string{1: "1", 2: "2"}
)
