package main

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
)

func main() {
	keywords := flag.String("keywords", "", "Keywords to search for.")
	flag.Parse()

	keywordsList := strings.Split(*keywords, " ")
	URL := "https://www.imdb.com/search/keyword/?keywords=" + keywordsList[0]
	for i := 1; i < len(keywordsList); i++ {
		URL += "%2C" + keywordsList[i]
	}

	c := colly.NewCollector()

	c.OnHTML(`h3[class="lister-item-header"]`, func(element *colly.HTMLElement) {
		fmt.Println(strings.TrimSpace(element.DOM.Children().Text()))
	})

	c.OnRequest(func(request *colly.Request) {
		log.Println("Visiting:", request.URL.String())
	})

	err := c.Visit(URL)
	if err != nil {
		log.Fatal(err)
	}
}
