package main

import (
	"fmt"
)

func main() {
	// non primitive datatypes slices and array are mutable data type , but array has a cauguage
	// primitive datatype string int are non mutable

	// slices
	// slice point the same reference in memory

	var slice []int = []int{1, 2, 3, 4}
	var slice1 = slice
	slice1[2] = 8
	fmt.Println(slice, slice1)
	fmt.Println("----------------------")

	// array
	// array copy from reference and create a new reference

	var array [3]int = [3]int{1, 2, 3}
	// array := [3]int{1, 2, 3}
	var array1 = array
	array[2] = 9

	fmt.Println(array, array1)
	fmt.Println("----------------------")

	// string int
	var num int = 1
	var num1 = num
	num = 3
	fmt.Println(num, num1)

}
