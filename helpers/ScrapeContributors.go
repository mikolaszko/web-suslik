package helpers

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func ScrapeContributors(args []string) {

	c := colly.NewCollector(
		colly.AllowedDomains("gitlab.com"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	c.OnHTML("span", func(e *colly.HTMLElement) {
		link := e.Text
		fmt.Println(link)
	})

	c.Visit(args[0])

}
