package main

import (
	"fmt"
)

func main() {

	msg := FunctionDeclaration()
	fmt.Println(msg)

	expressionMsg := FunctionExpression()
	fmt.Println(expressionMsg)

	var increment = ClosureFunction()
	fmt.Println(increment())
	fmt.Println(increment())
	fmt.Println(increment())
}

// function declaration
func FunctionDeclaration() string {
	return "FunctionDeclaration Invoked"
}

// function expression
var FunctionExpression = func() string {
	return "FunctionExpression Invoked"
}

// Closure Function
func ClosureFunction() func() int {
	x := 0
	return func() int {
		x += 1
		return x
	}
}
