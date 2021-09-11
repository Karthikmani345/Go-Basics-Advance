package main

import "fmt"

func main() {
	// var myName string  // delcaration
	// myName = "karthik"  // intialization
	// fmt.Println(myName)

	var myNumber int = 1 // instantiation
	fmt.Println(myNumber)

	const isCool bool = false
	// isCool = true; // cannot reassign the variable as its a constant
	fmt.Println(isCool)

	// Type inference
	myInferece := "let me interpret and guess the type"
	// myInferece = false
	fmt.Println(myInferece)

	// Mixed variable declaration / instantiation
	var firstName, lastName, age, isIndian = "karthik", "mani", 28, true
	fmt.Println(firstName, lastName, age, isIndian)

	var someVal1, someVal2 string
	someVal1 = "some value"
	someVal2 = "some value"
	fmt.Println(someVal1, someVal2)
}
