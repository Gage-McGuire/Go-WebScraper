package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

// site to be scraped
var url string = "https://books.toscrape.com"

func main() {
	c := colly.NewCollector(colly.AllowedDomains("books.toscrape.com"))

	var items []item

	c.OnHTML("article[class=product_pod]", func(e *colly.HTMLElement) {
		item := item{
			Name:   e.ChildText("h3"),
			Price:  e.ChildText("p[class=price_color]"),
			ImgUrl: e.ChildAttr("a", "href"),
		}

		items = append(items, item)
	})

	c.OnHTML("[class=next]", func(e *colly.HTMLElement) {
		next_page := e.Request.AbsoluteURL(e.ChildAttr("a", "href"))
		c.Visit(next_page)
	})

	c.OnRequest(func(e *colly.Request) {
		fmt.Println(e.URL.String())
	})

	c.Visit(url)

	file, err := os.Create("scraped.csv")
	if err != nil {
		log.Fatalln("Failed to create file", err)
	}
	defer file.Close()

}
