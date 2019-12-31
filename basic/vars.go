package basic

import "fmt"

// Outside a function, every construct begins with a keyword (var, func, and so on) and the := construct is not available

var (
	name            string
	location, email string
)

var (
	address                string = "address-1"
	age                    int    = 32
	addressInfer                  = "address-type-infer"
	multiLine1, multiLine2        = 1, "multiLine2"
)

// Basic run list
func Basic() {

	// Inside a function, the := can be used in place of a var declaration with implicit type.
	// Outside a function, := is not allowed
	typeAssignment := true

	fmt.Println(address)
	fmt.Printf("boolean value: %t\n", typeAssignment)
}
