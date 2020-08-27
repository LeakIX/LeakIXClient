package LeakIXClient

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
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
	SearchResults []SearchResult
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

func (sc *SearchResultsClient) SearchResult() SearchResult {
	return sc.SearchResults[sc.Position-1]
}

func GetSearchResults(scope string, query string, page int) ([]SearchResult, error) {
	url := fmt.Sprintf(
		"https://leakix.net/search?scope=%s&q=%s&page=%d", url2.QueryEscape(scope), url2.QueryEscape(query), page)
	var searchResults []SearchResult
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	resp, err := HttpClient.Do(req)
	if err != nil {
		return searchResults, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return searchResults, err
	}
	err = json.Unmarshal(body, &searchResults)
	if err != nil {
		return searchResults, err
	}
	return searchResults, nil
}

func GetChannel(scope string) (chan SearchResult, error) {
	channel := make(chan SearchResult)
	wsConnection, _, err := websocket.DefaultDialer.Dial("wss://leakix.net/ws/" + scope, map[string][]string{"Origin":{"leakix.net:443"}})
	if err != nil {
		return nil, err
	}
	go func() {
		searchResult := SearchResult{}
		for  {
			err := wsConnection.ReadJSON(&searchResult)
			if err != nil {
				log.Println("Error parsing websocket results. Is your scope correct?")
				log.Fatal(err)
			}
			channel <- searchResult
		}
	}()
	return channel, nil
}