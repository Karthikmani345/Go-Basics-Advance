package main

import (
	"fmt"
)

func main() {

	x := 1
	y := 2
	if x > y {
		fmt.Println(x, "greater than", y)
	} else {
		fmt.Println(x, "smaller than", y)
	}

	color := "blue"
	switch color {
	case "yellow":
		fmt.Println("its Yellow")
	case "red":
		fmt.Println("its red")
	default:
		fmt.Println("nothing")
	}
}
