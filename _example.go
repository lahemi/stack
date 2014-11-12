package main

import (
	"fmt"
	"github.com/lahemi/stack"
)

var (
	Istack = stack.Stack{}
	Sstack = stack.Stack{}
)

func main() {
	Istack.Push(1)
	Istack.Push(2)
	Istack.Push(3)

	i1, err := Istack.PopE()
	if err != nil {
		panic(err) // `emptyStackError`
	}
	// Or just ignore errors...
	i2 := Istack.Pop().(int)
	i3 := Istack.Pop().(int)
	i4 := Istack.Pop()

	fmt.Println(i1.(int) + i2 + i3)
	fmt.Printf("No values returns a nil: %v\n", i4)

	Sstack.Push("go")
	Sstack.Push("go")
	Sstack.Push("pher")
	fmt.Println(Sstack)
	Sstack.Swap() // You could check for an error.
	fmt.Println(Sstack)
	Sstack.Over()
	fmt.Println(Sstack)
	Sstack.Rot()
	fmt.Println(Sstack)

	// You can mix types too.
	Sstack.Push(5)
	fmt.Printf("%v %v\n", Sstack.Pop(), Sstack.Pop())

	// The only thing that matters when actually using
	// the values is doing the type assertions correctly.
}
