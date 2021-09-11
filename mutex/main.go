package main

import (
	"errors"
	"fmt"
	"sync"
)

// mutex is mutual exclusive lock

var (
	x     = 0
	mutex sync.Mutex
	w     sync.WaitGroup
)

func increment(wg *sync.WaitGroup) {
	mutex.Lock() // mutual exclusive lock
	x = x + 1
	mutex.Unlock() // release the lock
	wg.Done()
}
func main() {

	for i := 0; i < 100; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)

	err := errors.New("test error")
	fmt.Println(err)
}
