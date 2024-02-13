package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

// site to be scraped
var url string = "https://books.toscrape.com"

// scrape selector
func main() {
	scrape(url)
}

func scrape(url string) {
	c := colly.NewCollector(colly.AllowedDomains("books.toscrape.com"))

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Visting", request.URL)
	})

	c.OnHTML("article[class=product_pod]", func(e *colly.HTMLElement) {
		item := item{
			Name:   e.ChildText("h3"),
			Price:  e.ChildText("p[class=price_color]"),
			ImgUrl: e.ChildAttr("a", "href"),
		}

		fmt.Println(item)
	})

	c.Visit(url)
}
