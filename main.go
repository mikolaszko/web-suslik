package main

import (
	"log"

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
	susCmd = &cobra.Command{
		Use:   "sus [link to scrape]",
		Short: "Scrapes provided link for hrefs",
		Args:  cobra.MaximumNArgs(1),
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
	rootCmd.AddCommand(susCmd)
	// echoCmd.AddCommand(timesCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
