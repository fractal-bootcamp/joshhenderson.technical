/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "scraper",
	Short: "a web scraper application",
	Long:  `Takes a URL, the number of links per page, and the depth it should crawl and returns a parsed object at the CL`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Executing web scraping...")
		resp, err := http.Get("https://en.wikipedia.org/wiki/Main_Page") // get to wikipedia main page
		if err != nil {
			log.Fatal(err) // handle error
		}
		defer resp.Body.Close()

		bytes, _ := io.ReadAll(resp.Body)

		// print to console
		fmt.Println("Response status:", resp.Status)
		fmt.Println(string(bytes))
		os.WriteFile("websites/test.html", bytes)
		// fetch the page
		// clean the page
		// save the page
		// get the links
		// go through each link and do the same

		// file, _ := os.Open("./cmd/example.txt")
		// bytes, _ := io.ReadAll(file)

		// fmt.Println(bytes)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.scraper.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//grootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
