package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

func predictFunc(ctx context.Context, dataCh chan int64) (int64, error) {
	start := time.Now()
	defer func() {
		fmt.Println(time.Since(start))
	}()

	go func() {
		select {
		case <-ctx.Done():
		case dataCh <- unpredictableFunc():
		}
	}()

	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case val := <-dataCh:
		return val, nil
	}

}

func main_predfunc() {
	timeout := 400 * time.Millisecond
	wg := sync.WaitGroup{}

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	dataCh := make(chan int64)

	wg.Add(1)
	go func() {
		defer wg.Done()
		val, err := predictFunc(ctx, dataCh)
		fmt.Println(val, err)
	}()

	wg.Wait()

}
