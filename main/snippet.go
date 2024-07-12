package main

import (
	"regexp"
	"sync"
	"time"
)

// itemDescription симулирует получение описания из сервиса
func itemDescription(itemID int) string {
	time.Sleep(time.Second * 1)
	return "☆☆☆Lorem ♡♡♡ ipsum  dolor sit amet..."
}

// itemPrice симулирует получение цены в долларах из сервиса
func itemPrice(itemId int) float64 {
	time.Sleep(2 * time.Second)
	return 100
}

func prettify(description string) string {
	time.Sleep(time.Second * 1)
	re := regexp.MustCompile("\\W+")
	return re.ReplaceAllString(description, "")
}

func priceToRub(price float64) float64 {
	time.Sleep(1 * time.Second)
	return price * 70
}

type Snippet struct {
	Price       float64
	Description string
}

func BuildSnippet(itemId int) Snippet {
	var wg sync.WaitGroup

	var desc string

	wg.Add(2)
	go func() {
		desc = prettify(itemDescription(itemId))
		wg.Done()

	}()

	var price float64
	go func() {
		price = priceToRub(itemPrice(itemId))
		wg.Done()
	}()

	wg.Wait()

	return Snippet{
		Price:       price,
		Description: desc,
	}

}
