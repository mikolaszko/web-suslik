package helpers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeHeadlines(link []string) {

	rootDomain :=
		strings.Split(strings.SplitAfter(link[0], ".com")[0], "//")[1]
	c := colly.NewCollector(
		colly.AllowedDomains(rootDomain),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		if e == nil {
			fmt.Println("No headings found on this page")
		}
		fmt.Println(e.Text)
	})

	c.Visit(link[0])

}
