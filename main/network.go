package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequest int = 1000

var count int32

func networkRequest() {
	time.Sleep(time.Millisecond)
	atomic.AddInt32(&count, 1)
}

func main_network() {
	// var max int
	// var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < numRequest; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			networkRequest()
		}()
	}
	wg.Wait()
	fmt.Println(count)

}
