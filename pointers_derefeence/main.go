package main

import (
	"fmt"
)

func main() {
	//  & pointers and  * dereference

	var a int = 1
	fmt.Println("pointer", &a)
	fmt.Println("value", a)

	// storing the address of the variable by address operator (&) aka pointer
	// The & operator generates a pointer to its operand.
	var b = &a
	fmt.Println("pointer b", b)
	//	The * operator denotes the pointer's underlying value.
	*b = 2
	// fmt.Println("value b", &b)

	fmt.Println(a, *b)
	fmt.Println("----------------------")

	var x int = 1
	var y = &x
	fmt.Println(x, y)
	// deference the variable by prepending star in variable
	// using * operator before a pointer
	// variable to access the value stored at the variable at which it is pointing
	*y = 2
	fmt.Println(x, y)
	fmt.Println("----------------------")

	var z int = 3
	var t = changePointer(&z)
	fmt.Println(*t)
	fmt.Println("----------------------")

}

func changePointer(pointer *int) *int {
	*pointer = 2
	return pointer
}
