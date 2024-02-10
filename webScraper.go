package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

var url string = "https://www.amazon.com/Gopher-Lover-Pocket-T-Shirt/dp/B09FZQ55K1"

type Product struct {
	name string
}

func main() {
	scrape(url)
}

func scrape(url string) {
	c := colly.NewCollector()

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting", request.URL)
	})

	c.OnHTML("#dp-container", func(e *colly.HTMLElement) {
		name := e.ChildText("#centerCol #productTitle")
		fmt.Println(Product{name})
	})

	c.Visit(url)
}
