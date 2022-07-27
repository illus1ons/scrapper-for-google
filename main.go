package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"
	"math/rand"
	"net/url"
	"github.com/PuerkitoBio/goquery"
)

var googleDomains = map[string]string{
	"com": "https://www.google.com/search?q=",
	"kr": "https://www.google.co.kr/search?q=",
}

type SearchResult struct {
	ResultRank int
	ResultURL string
	ResultTitle string
	ResultDesc string
}

var userAgents = []string{

}

func randomUserAgent() string {
	randNum := rand.Int() % len(userAgents)
	return userAgents[randNum]
}

func buildGoogleUrls(searchTerm, countryCode, pages, count int) () {
	toScrape := []string{}
	searchTerm = strings.Trim(searchTerm, " ")
	searchTerm = strings.Replace(searchTerm, " ", "+", -1)
	if googleBase, found := googleDomains[countryCode]; found {
		for i := 0, i < pages ; i++ {
			start := i * count
			scrapeURL := fmt.Sprints(googleBase, searchTerm, count, start)
		}
	}
}

func GoogleScrape(searchTerm, countryCode, pages, count) ([]SearchResult, err) {
	results := []SearchResult{}
	resultCounter := 0
	googlePages, err := buildGoogleUrls(searchTerm, countryCode, pages, count)
}

func main() {
	res, err := GoogleScrape("react next.js", "kr", 1, 30)
	if err == nil {
		for _, res := range res {
			fmt.Println(res)
		}
	}
}