package helpers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeHeadlines(args []string) {

	rootDomain :=
		strings.Split(strings.SplitAfter(args[1], args[0])[0], "//")[1]

	c := colly.NewCollector(
		colly.AllowedDomains(rootDomain),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	c.OnHTML("h1", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})

	c.OnHTML("h2", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML("h3", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML("h4", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.OnHTML("h5", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
	})
	c.Visit(args[1])

}
