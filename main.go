package main

import (
	"fmt"
	"go-playground/basic"
	"go-playground/basic/collections"
	"sync"
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

	res11 := collections.IntArray
	// use %q to print each element as equivalent quoted string
	fmt.Printf("%q\n", res11)
	fmt.Printf("%d\n", res11)
	// same as fmt.Printf("%d\n", res11)
	fmt.Println(res11)

	res12 := collections.MultiDimensionArray
	for i := 0; i < 2; i++ {
		for j := 0; j < 1; j++ {
			res12[i][j] = fmt.Sprintf("row %d - column %d", i, j)
		}
	}
	fmt.Printf("%q\n", res12)
	fmt.Println(res12)

	// Chp. 7 interfaces
	res13 := basic.Greet(&basic.User{FirstName: "firstName", LastName: "lastName"})
	fmt.Println(res13)

	var wg sync.WaitGroup

	// Chp. concurrency
	//wg.Add(2)

	wg.Add(1)
	go basic.GoroutinePrint(&wg, "Concurrency!")

	wg.Add(1)
	go basic.GoroutinePrintf(&wg, "Concurrency Print Evaluate")

	// we cannot do as follow since goroutine only pass value via channel
	// so not sure when the return result will pass as value to another func
	// fmt.Println(basic.GoroutinePrintf(&wg, "Concurrency Print Evaludate first"))
	//wg.Wait()

	ch := make(chan int)
	wg.Add(2)
	go basic.ChannelPass(5, ch, &wg)
	go basic.ChannelPass(15, ch, &wg)
	// not guarantee order, but if we combine defer wg.Done()

	fmt.Printf("%d, %d \n", <-ch, <-ch)
	// wait done should wait channel calls are all done.
	wg.Wait()

	chFib := make(chan int, 10)
	go basic.ChannelFib(cap(chFib), chFib)
	basic.ChannelRange(chFib)

	fmt.Println("channel select:")
	// select case with go channel
	chSelect := make(chan int)
	chSelectQuit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-chSelect)
		}
		chSelectQuit <- 0
	}()
	basic.ChannelSelect(chSelect, chSelectQuit)
}
