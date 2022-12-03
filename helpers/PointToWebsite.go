package helpers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func PointToWebsite(link []string) {

	rootDomain :=
		strings.Split(strings.SplitAfter(link[0], ".com")[0], "//")[1]

	c := colly.NewCollector(
		colly.AllowedDomains(rootDomain),
	)

	c.OnHTML(".spacer", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})

	c.Visit(link[0])

}
