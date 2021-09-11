package main

import (
	"fmt"
	"time"
)

//  select statement will block until one of its cases is executed.

func main() {
	promise1 := make(chan bool)
	defer close(promise1)

	promise2 := make(chan bool)
	defer close(promise2)

	go MockPromise(3, promise1)
	go MockPromise(5, promise2)

	// var _ = <-promise1
	// var _ = <-promise2

	// select is like race condition i.e promise.race()
	select {
	case <-promise1:
		fmt.Println("promise 1 resolved")
	case <-promise2:
		fmt.Println("promise 2 resolved")
		// default:
		// 	fmt.Println("default case executed")
	}
}

func MockPromise(threshold int, done chan bool) {
	time.Sleep(time.Second * time.Duration(threshold))
	// fmt.Println("channel with threshold:", threshold, " resolved")
	done <- true
}
