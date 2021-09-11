package main

import "fmt"

func main() {

	// declaring the array
	var myArray [3]string
	myArray[0] = "karthik"
	myArray[1] = "mani"

	// fmt.Println(myArray)
	//fmt.Println(len(myArray))
	fmt.Println(myArray[1])
	fmt.Println(myArray[2])

	// intializing the arrays
	var initArray = []string{"karthik", "mani"}
	fmt.Println(initArray)

	initInteger := [3]int{
		1, 2, 3,
	}
	fmt.Println(initInteger)

	// // Range (is like foreach)
	for index, element := range initArray {
		fmt.Println("index", index, "element", element)
	}

}
