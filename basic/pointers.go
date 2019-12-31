package basic

import "fmt"

// Go has pointers, but no pointer arithmetic

// Only constant is immutable
// and most case are pass by value
// for pass by reference, use pointer

// PointerAsReference (y *int, z *int)
func PointerAsReference(y *int, z *int) *int {
	*y = (*y + *z)
	fmt.Println(*y)
	fmt.Println(y)

	fmt.Println(*z)
	fmt.Println(z)
	return y
}
