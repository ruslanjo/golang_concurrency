package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var weather int
var mu sync.RWMutex

// aiWeatherForecast через нейронную сеть вычисляет температуру за ~1 секунду
func aiWeatherForecast() int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func updateWeather() {
	forecast := aiWeatherForecast()
	mu.Lock()
	weather = forecast
	mu.Unlock()
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	updateWeather()

	go func(ctx context.Context) {
		ticker := time.NewTicker(10 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				updateWeather()
			case <-ctx.Done():
				return
			}
		}
	}(ctx)


	http.HandleFunc("/weather", getWeatherForecast())
	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}

func getWeatherForecast() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		fmt.Fprintf(w, "{\"temperature\":%d}\n", weather)
		mu.RUnlock()
	}
}
