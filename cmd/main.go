package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

const baseUrlFmt = "https://reviews.birdeye.com/%s"

func main() {
	companyId := flag.String("company-id", "", "a valid reviews.birdeye.com company id")
	flag.Parse()
	if companyId == nil || *companyId == "" {
		os.Exit(1)
	}

	site := fmt.Sprintf(baseUrlFmt, *companyId)
	log.WithField("site", site).Info("Scraping website")
}
