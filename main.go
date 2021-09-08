package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML(`h3[class="lister-item-header"]`, func(element *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(element.Text))
	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting:", request.URL.String())
	})

	err := c.Visit("https://www.imdb.com/search/keyword/?keywords=investigation")
	if err != nil {
		log.Fatal(err)
	}
}
