package main

import (
	"encoding/json"
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

	doc, err := NewDoc(site)
	if err != nil {
		fmt.Printf("Error getting site: %s\n", err.Error())
		os.Exit(1)
	}

	i := 0
	revs := []Review{}
	for {
		fmt.Printf("Scraping page: %d\n", i)
		i++

		doc.Find("div.Review__contentWrapper__2NQN3").Each(
			func(index int, selector *goquery.Selection) {
				source, ok := selector.Find("a").Attr("href")
				if !ok {
					return
				}

				rev := Review{
					Source: source,
					Date:   selector.Find("div.__react_component_tooltip").Text(),
					Stars:  len(selector.Find("span.RatingStar__be-star-on__28Wmg").Nodes),
					Text:   selector.Find("span.Review__reviewPara__2qFYA").Text(),
				}
				revs = append(revs, rev)
			})

		next := doc.Find("li.next")
		if next.HasClass("disabled") {
			break
		}

		page, ok := next.Find("a").Attr("href")
		if !ok {
			break
		}

		doc, err = NewDoc(fmt.Sprintf(baseUrlFmt, page))
		if err != nil {
			fmt.Printf("Error getting site: %s\n", err.Error())
			os.Exit(1)
		}
	}

	json, err := json.MarshalIndent(revs, "", "  ")
	if err != nil {
		fmt.Printf("Error marshalling json: %s\n", err.Error())
		os.Exit(1)
	}

	os.WriteFile("reviews.json", json, 0644)
}

func NewDoc(site string) (*goquery.Document, error) {
	res, err := http.Get(site)
	if err != nil || res.StatusCode != http.StatusOK {
		fmt.Printf("Error getting website: %s\n", err.Error())
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Printf("Error parsing response body: %s\n", err.Error())
		return nil, err
	}

	return doc, nil
}

type Review struct {
	Source string
	Date   string
	Stars  int
	Text   string
}
