/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/cobra"
)

func writeFile(data, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.WriteString(data)
}

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

		// bytes, _ := io.ReadAll(resp.Body)
		//stringByte := string(bytes)
		// print to console
		fmt.Println("Response status:", resp.Status) //print reponse status
		//fmt.Println(string(bytes))                      //print the html tree
		// os.WriteFile("websites/test.html", bytes, 0777) //write the data to a file with open permissions
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(doc)

		h2 := doc.Find("h2").First().Text() //"returns text from first h2 tag "

		// section, err := doc.Find("div.mw-content-ltr").Html()
		// if err != nil {
		// 	log.Fatal(err)
		// }

		doc.Find("div.mw-content-ltr").Find("p").Each(func(index int, item *goquery.Selection) {
			url, _ := item.Find("a").Attr("href")
			fmt.Println(url)
		})

		//fmt.Println(links) // print links
		//writeFile(section, "wiki")
		//fmt.Println(section)
		fmt.Println(h2)
		fmt.Println("visual break")

		// file, _ := os.Open("./cmd/example.txt")
		// bytes, _ := io.ReadAll(file)
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

type Link struct {
	number, url string
}

var wikilinks []Link
