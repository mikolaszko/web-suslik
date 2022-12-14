package main

import (
	"log"

	"github.com/mikolaszko/web-suslik/helpers"
	"github.com/mikolaszko/web-suslik/encodingHelpers"
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
	SuslinksCmd = &cobra.Command{
		Use:   "sl [url to scrape, domain]",
		Short: "Scrapes provided https://url for hrefs",
		Args:  cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			helpers.ScrapeLinks(args)
		},
	}
	SusheadlinesCmd = &cobra.Command{
		Use: "sh [url to scrape]",
		Short: "Scrapes provided https://url for h1s",
		Args: cobra.MaximumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			helpers.ScrapeHeadlines(args)
		},
	}
	SusWikipediaContents = &cobra.Command{
		Use: "swc [url to scrape]",
		Short: "Scrapes provided wikipedia article and creates .csv file",
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			encodingHelpers.ScrapeWikipediaTableContent(args)
		},
	}
)

func init() {
	rootCmd.PersistentFlags().BoolVarP(&persistRootFlag, "persistFlat", "p", false, "a persistant root flag")
	rootCmd.Flags().BoolVarP(&localRootFlag, "localFlag", "l", false, "a local root flag")
	// timesCmd.Flags().IntVarP(&times, "times", "t", 1, "number of times to echo to stdout")
	rootCmd.AddCommand(SuslinksCmd)
	rootCmd.AddCommand(SusheadlinesCmd)	
	rootCmd.AddCommand(SusWikipediaContents)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
