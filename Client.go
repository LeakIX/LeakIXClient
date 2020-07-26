package LeakIXClient

import (
	"encoding/json"
	"fmt"
	"gitlab.nobody.run/tbi/core"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"time"
)

var LeakIXProxy = &core.ProxiedPlugin{}
var LeakIXHttpTranport = &http.Transport{
	DialContext:           LeakIXProxy.DialContext,
	ResponseHeaderTimeout: 5 * time.Second,
	ExpectContinueTimeout: 5 * time.Second,
}
var HttpClient = &http.Client{
	Transport: LeakIXHttpTranport,
	Timeout:   5 * time.Second,
}

type SearchResultsClient struct {
	Scope         string
	Query         string
	SearchResults []*SearchResult
	Position      int
	Page          int
}

func (sc *SearchResultsClient) Next() bool {
	if len(sc.SearchResults) > sc.Position {
		sc.Position++
		return true
	}
	// Try to load next page
	results, _ := GetSearchResults(sc.Scope, sc.Query, sc.Page)
	for _, result := range results {
		sc.SearchResults = append(sc.SearchResults, result)
	}
	sc.Page++
	if len(sc.SearchResults) > sc.Position {
		sc.Position++
		return true
	}
	return false
}

func (sc *SearchResultsClient) SearchResult() *SearchResult {
	return sc.SearchResults[sc.Position-1]
}

func GetSearchResults(scope string, query string, page int) ([]*SearchResult, error) {
	url := fmt.Sprintf(
		"https://leakix.net/search?scope=%s&q=%s&page=%d", url2.QueryEscape(scope), url2.QueryEscape(query), page)
	var searchResults []*SearchResult
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	resp, err := HttpClient.Do(req)
	if err != nil {
		log.Println(err)
		return searchResults, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return searchResults, err
	}
	err = json.Unmarshal(body, &searchResults)
	if err != nil {
		log.Println(err)
		return searchResults, err
	}
	return searchResults, nil
}
