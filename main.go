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

var googleDomains = map[string]string{}

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

func main() {
	res, err := GoogleScrape("akhil Sharma")
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}
