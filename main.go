package main

import (
	"fmt"
	"github.com/Puerkitobio/goquery"
	"math/rand"
	"net/http"
	"net/url"
	"strings"
	"time"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"za":  "https://www.google.co.za/search?q=",
}

type SearchResult struct {
	ResultRand  int
	ResultUrl   string
	ResultTitle string
	ResultDesc  string
}

var userAgents = []string{}

func randomUserAgent() string {
	randNum := rand.Intn(len(userAgents))
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, pages, count int)([]string, error) {
	toScrape := []string{}
	searchTerm := strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0, i < pages; i++ {
			start:= i * count
			scrapeUrl := fmt.Sprintf("%s%s&num=%d&hl=%s&start=%d&filter=0",googleBase, searchTerm, count, languageCode, start )
		}
	}else{
		err:= fmt.Errorf("country (%s) code not found in google", countryCode)
		return nil, err
	}
	return toScrape, nil
}

func GoogleScrape(searchTerm, countryCode, languageCode string, pages, count) ([]SearchResult, error) {
	result := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, languageCode, pages, count)
	if err != nil {
		return nil, err
	}
}

func main() {
	res, err := GoogleScrape("akhil Sharma", "com", "en", 1, 30)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
