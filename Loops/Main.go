package main

import "fmt"

func main() {
	// var i int
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// break
	for j := 0; j < 10; j++ {
		if j == 5 {
			fmt.Println("breaking statement")
			break
		}
		fmt.Println("iteration ", j)
	}

	// continue statement
	for x := 0; x < 10; x++ {
		if x == 5 {
			fmt.Println("continue statement")
			continue
		}
		fmt.Println("iteration ", x)
	}
}
