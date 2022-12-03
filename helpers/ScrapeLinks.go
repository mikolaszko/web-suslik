package helpers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeLinks(args []string) {

	rootDomain :=
		strings.Split(strings.SplitAfter(args[1], args[0])[0], "//")[1]

	c := colly.NewCollector(
		colly.AllowedDomains(rootDomain),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Println(link)
	})

	c.Visit(args[1])

}
