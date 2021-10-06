package LeakIXClient

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/LeakIX/l9format"
	"github.com/gorilla/websocket"
	"gitlab.nobody.run/tbi/core"
	"io/ioutil"
	"log"
	"net/http"
	url2 "net/url"
	"strings"
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
	SearchResults []l9format.L9Event
	Position      int
	Page          int
	ApiKey        string
	Endpoint      string
	LastError     error
}

const defaultEndpoint = "https://leakix.net"

func (sc *SearchResultsClient) GetEndpoint() string {
	if len(sc.Endpoint) > 8 {
		return sc.Endpoint
	}
	return defaultEndpoint
}

func (sc *SearchResultsClient) Next() bool {
	var results []l9format.L9Event
	if len(sc.SearchResults) > sc.Position {
		sc.Position++
		return true
	}
	// Try to load next page
	results, sc.LastError = sc.GetSearchResults(sc.Scope, sc.Query, sc.Page)
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

func (sc *SearchResultsClient) SearchResult() l9format.L9Event {
	return sc.SearchResults[sc.Position-1]
}

func (sc *SearchResultsClient) GetSearchResults(scope string, query string, page int) ([]l9format.L9Event, error) {
	url := fmt.Sprintf(
		"%s/search?scope=%s&q=%s&page=%d", sc.GetEndpoint(), url2.QueryEscape(scope), url2.QueryEscape(query), page)
	var searchResults []l9format.L9Event
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("api-key", sc.ApiKey)
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

func (sc *SearchResultsClient) GetChannel(scope string) (chan l9format.L9Event, error) {
	channel := make(chan l9format.L9Event)
	endpointUrl, err := url2.Parse(sc.GetEndpoint())
	if err != nil {
		return nil, errors.New("invalid endpoint")
	}
	endpointUrl.Scheme = strings.Replace(endpointUrl.Scheme, "http", "ws", -1)
	log.Println(endpointUrl.String())
	wsConnection, _, err := websocket.DefaultDialer.Dial(endpointUrl.String()+"/ws/"+scope, map[string][]string{
		"Origin": {endpointUrl.Host + ":" + endpointUrl.Port()},
		"api-key":{sc.ApiKey},
	})
	if err != nil {
		return nil, err
	}
	go func() {
		searchResult := l9format.L9Event{}
		for {
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
