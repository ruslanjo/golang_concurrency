package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func main_urls() {
	var wg sync.WaitGroup
	hosts := []string{
		"https://google.com",
		"https://avito.ru",
	}

	for _, uri := range hosts {
		wg.Add(1)
		go func(uri string) {
			defer wg.Done()
			makeRequest(uri)
		}(uri)
	}
	wg.Wait()
}

type requestResult struct {
	Uri        string
	StatusCode int
	Err        string
}

func makeRequest(uri string) {
	fmt.Println(uri)
	resp, err := http.Get(uri)
	reqResult := requestResult{
		Uri:        uri,
		StatusCode: resp.StatusCode,
	}

	if err != nil {
		reqResult.Err = err.Error()
	}
	jsonData, _ := json.Marshal(reqResult)
	fmt.Println(string(jsonData))
}
