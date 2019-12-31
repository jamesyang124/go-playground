package main

import (
	"fmt"
	"playground/basic"
)

func main() {
	fmt.Printf("hello, world\n")

	a := [...]int{1, 2, 3, 4, 5}
	b := [3]int{3, 2, 1}
	var c [5]int
	d := [5]int{1: 1, 3: 3}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	for i := 5; i < 15; i++ {
		fmt.Println(i)
	}

	basic.Basic()

	typeInference1 := basic.Constant2
	fmt.Println(typeInference1)

	res1, res2 := basic.ReturnMultiple(5, 6)
	fmt.Printf("res1: %d, res2: %d\n", res1, res2)

	res3, res4 := basic.ReturnNamed(5, 6)
	fmt.Printf("res3: %d, res4: %d\n", res3, res4)

	res5 := basic.PointerAsReference(&res3, &res4)
	fmt.Printf("res5 pointer: %p, de-reference value: %d\n", res5, *res5)

	var unInitVarZ interface{}
	fmt.Println(basic.TypeAssertion(unInitVarZ))

	res6 := basic.BootcampPtr
	fmt.Println(*res6)

	res7 := basic.InitBootcamp
	fmt.Println(*res7)
	// pointer of new Bootcamp
	fmt.Println(res7)

	res8 := &basic.Bootcamp{}
	// init data the same with other ptr have same value
	fmt.Println(*res8 == *res7)
	// but ptr not the same
	fmt.Println(res8 == res7)
	fmt.Println(res8)

	res9 := basic.ImplicitCompose
	fmt.Println(res9)

	res10 := basic.ExplicitCompose
	// res10.composition undefined (cannot refer to unexported field or method composition)
	// res10.composition = basic.Composition{val: 41, Text: "explicit"}
	// this due to composition is non-public field so inaccessible
	//
	// res10.Compose = basic.Composition{val: 41, Text: "explicit"}
	// same here, val: 41 in inaccessible
	// unknown field 'val' in struct literal of type basic.Composition
	res10.Compose = basic.Composition{Text: "explicit"}

	fmt.Println(res10)
}
