package main

import (
	"fmt"
	"sync"
)

func main_cycle() {
	var max int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1000; i > 0; i-- {
		i := i
		wg.Add(1)

		go func() {
			defer wg.Done()
			if i%2 != 0{
				return
			}
			mu.Lock()
			if i > max {
				max = i
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(max)
}
