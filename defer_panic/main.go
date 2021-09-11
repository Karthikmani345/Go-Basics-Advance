package main

import (
	"fmt"
)

// *** fundamentals ***
// panic : try or throw error
// recover : catch block
// defer : finally

func main() {
	// recover can be called only inside defer function
	defer CustomRecover()
	fmt.Println("main function started")
	defer fmt.Println("main function ended")
	Run()
}

func Run() {
	// recover can be called only inside defer function
	// defer CustomRecover()
	defer fmt.Println("Run completed")
	fmt.Println("Run Started")
	// panic("Panic error and stoping execution")
	name := "karthik"
	fmt.Println("printing variable", name)
}

func CustomRecover() {
	r := recover()
	fmt.Println("recover() :", r)
	if r != nil {
		fmt.Println("recovered :", r)
	}
}
