package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(`h3[class="lister-item-header"]`, func(element *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(element.DOM.Children().Text()))
	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting:", request.URL.String())
	})

	err := c.Visit("https://www.imdb.com/search/keyword/?keywords=investigation")
	if err != nil {
		log.Fatal(err)
	}
}
