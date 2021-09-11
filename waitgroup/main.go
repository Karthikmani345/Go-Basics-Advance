package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Worker start")
		time.Sleep(time.Second)
		fmt.Println("Worker done")
	}()

	wg.Wait()
	fmt.Println("------------")

}
