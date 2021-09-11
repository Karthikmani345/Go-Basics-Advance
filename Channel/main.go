package main

import (
	"fmt"
	"time"
)

func main() {
	channel := mockIntensiveTask()
	fmt.Println("waiting to get data from channel")
	var result = <-channel
	if result.err == nil {
		fmt.Println("result data", result.data)
	} else {
		fmt.Println("error", result.err)
	}

}

func mockIntensiveTask() chan result {
	// create promise by making a channel
	fmt.Println("channel i.e promise created")
	channel := make(chan result)

	// defering the execution i.e asycronous using goroutines
	go func() {
		defer close(channel)
		time.Sleep(time.Second * 3)
		fmt.Println("channel i.e promise resolved")
		channel <- result{
			err:  nil,
			data: true,
		}
	}()

	return channel
}

type result struct {
	err  error
	data interface{} // any type
}
