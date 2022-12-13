package helpers

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func ScrapeContributors(args []string) {

	c := colly.NewCollector(
		colly.AllowedDomains("github.com"),
	)

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		if e.Attr("class") == "avatar avatar-user" {
		link := e.Attr("src")
		fmt.Println(link)
		}
	})
	builder := strings.Builder{}
        builder.WriteString(args[0])
 	builder.WriteString("/graphs/contributors")
	fmt.Println(builder.String())

	c.Visit(builder.String())

}
