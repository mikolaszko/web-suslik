package encodingHelpers 

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly/v2"
)

func ScrapeWikipediaTableContent(args []string) {
	fileName := "data.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create file, err: %q", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)
	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
	})
	c.OnHTML("li.toclevel-1", func(e *colly.HTMLElement){
		e.ForEach("a", func(_ int, el *colly.HTMLElement) {
			writer.Write(([]string{
				el.ChildText("span.tocnumber"),
				el.ChildText("span.toctext"),
			}),	
		)},
	)},
	)
	fmt.Println("Scraping is finished")
	c.Visit(args[0])
}
