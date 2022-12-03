package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly/v2"
	"github.com/mikolaszko/web-suslik/helpers"
	"github.com/spf13/cobra"
)

var (
	localRootFlag   bool
	persistRootFlag bool
	// times           int
	rootCmd = &cobra.Command{
		Use:   "example",
		Short: "An example cobra program",
		Long:  "A longer example cobra program",
	}
	echoCmd = &cobra.Command{
		Use:   "echo [strings to echo]",
		Short: "Prints given strings to stdout",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			helpers.PointToWebsite(args)
		},
	}
	// timesCmd = &cobra.Command{
	// 	Use:   "times [strings to echo]",
	// 	Short: "prints given strings to stdout multiple times",
	// 	Args:  cobra.MinimumNArgs(1),
	// 	Run: func(cmd *cobra.Command, args []string) {
	// 		for i := 0; i < times; i++ {
	// 			fmt.Println("Echo: " + strings.Join(args, " "))
	// 		}
	// 	},
	// }
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlat", "p", false, "a persistant root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	// timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	rootCmd.AddCommand(echoCmd)
	// echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
	c := colly.NewCollector(
		colly.AllowedDomains("en.wikipedia.org"),
	)

	c.OnHTML(".mw-parser-output", func(e *colly.HTMLElement) {
		links := e.ChildAttrs("a", "href")
		fmt.Println(links)
	})
	c.Visit("https://en.wikipedia.org/wiki/Web_scraping")
}
