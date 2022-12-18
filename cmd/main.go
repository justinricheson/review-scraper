package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const baseUrlFmt = "https://reviews.birdeye.com/%s"

func main() {
	companyId := flag.String("company-id", "", "a valid reviews.birdeye.com company id")
	flag.Parse()
	if companyId == nil || *companyId == "" {
		fmt.Println("Error parsing company-id flag")
		os.Exit(1)
	}

	site := fmt.Sprintf(baseUrlFmt, *companyId)
	fmt.Printf("Scraping website: %s", site)

	res, err := http.Get(site)
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Printf("Error getting website: %s\n", err.Error())
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Printf("Error parsing response body: %s\n", err.Error())
		os.Exit(1)
	}

	doc.Find("div.Review__contentWrapper__2NQN3").Each(
		func(index int, selector *goquery.Selection) {
			review := selector.Find("span.Review__reviewPara__2qFYA").Text()
			fmt.Printf("Found review: %s\n", review)
		})
}
