package main

import "fmt"

func main() {
	var myName string = "karthik"
	var myNumber int = 1
	var myNumber32 int32 = 1
	var myBoolean bool = false

	fmt.Println(myName)
	fmt.Println(myNumber)
	fmt.Println(myBoolean)
	fmt.Println(myNumber32)

	// declared a variable without assigning it prints the default value
	var testNum uint16
	fmt.Println(testNum)
}
