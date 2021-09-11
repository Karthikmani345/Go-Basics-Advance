package main

import (
	"fmt"
)

func main() {
	// declare map
	var myDeclMap map[string]string
	myDeclMap = make(map[string]string)
	// myDeclMap["test"] = "value"
	fmt.Println(myDeclMap)

	// define a map using make i.e make intialize the value with nil value
	var myMap = make(map[string]string)
	myMap["name"] = "karthik"
	fmt.Println(myMap)

	// intantiaztion
	var myMapInit = map[string]string{
		"name":     "karthik mani",
		"lastName": "subbhaya",
		"petName":  "SP",
	}

	// delete from a map
	delete(myMapInit, "lastName")
	fmt.Println(myMapInit)

	// loopin through the map object
	for key, value := range myMapInit {
		fmt.Println("key", key, "value", value)
	}

}
